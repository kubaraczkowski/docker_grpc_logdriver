package driver

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"syscall"

	"github.com/containerd/fifo"
	"github.com/docker/docker/api/types/plugins/logdriver"
	"github.com/docker/docker/daemon/logger"
	protoio "github.com/gogo/protobuf/io"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type driver struct {
	mu     sync.Mutex
	logs   map[string]*logPair
	idx    map[string]*logPair
	recvCh chan *LogMessage
	stopCh chan struct{}
	UnimplementedIDockerLogDriverServer
}

type logPair struct {
	stream io.ReadCloser
	info   logger.Info
	sendCh chan *LogMessage
	stopCh chan struct{}
	msgCh  chan logdriver.LogEntry
}

type fanout struct {
	mu            sync.RWMutex
	clientchannel map[uuid.UUID](chan *LogMessage)
}

var clientfanout fanout

func NewDriver() *driver {
	return &driver{
		logs:   make(map[string]*logPair),
		idx:    make(map[string]*logPair),
		recvCh: make(chan *LogMessage),
		stopCh: make(chan struct{}),
	}
}

func (d *driver) StartLogging(file string, logCtx logger.Info) error {
	d.mu.Lock()
	if _, exists := d.logs[file]; exists {
		d.mu.Unlock()
		return fmt.Errorf("logger for %q already exists", file)
	}
	d.mu.Unlock()

	logrus.WithField("id", logCtx.ContainerID).WithField("file", file).WithField("logpath", logCtx.LogPath).Debugf("Start logging")
	f, err := fifo.OpenFifo(context.Background(), file, syscall.O_RDONLY, 0700)
	if err != nil {
		return errors.Wrapf(err, "error opening logger fifo: %q", file)
	}

	d.mu.Lock()
	lf := &logPair{
		stream: f,
		info:   logCtx,
		sendCh: d.recvCh,
		stopCh: d.stopCh,
		msgCh:  make(chan logdriver.LogEntry),
	}
	d.logs[file] = lf
	d.idx[logCtx.ContainerID] = lf
	d.mu.Unlock()

	go readLog(lf)
	go consumeLog(lf)
	return nil
}

func (d *driver) StopLogging(file string) error {
	logrus.WithField("file", file).Debugf("Stop logging")
	d.mu.Lock()
	lf, ok := d.logs[file]
	if ok {
		lf.stream.Close()
		delete(d.logs, file)
	}
	if len(d.logs) == 0 {
		close(d.stopCh)
	}
	d.mu.Unlock()
	return nil
}

func readLog(lf *logPair) {
	dec := protoio.NewUint32DelimitedReader(lf.stream, binary.BigEndian, 1e6)
	defer dec.Close()
	var buf logdriver.LogEntry
	for {
		if err := dec.ReadMsg(&buf); err != nil {
			if err == io.EOF {
				logrus.WithField("id", lf.info.ContainerID).WithError(err).Debug("shutting down log logger")
				lf.stream.Close()
				return
			}
			dec = protoio.NewUint32DelimitedReader(lf.stream, binary.BigEndian, 1e6)
			continue
		}
		lf.msgCh <- buf
	}
}

func consumeLog(lf *logPair) {

	for {

		buf := <-lf.msgCh

		msg_proto := &LogMessage{
			Service: lf.info.ContainerName,
			Entry: &LogEntry{
				Source:   buf.GetSource(),
				TimeNano: buf.GetTimeNano(),
				Line:     buf.GetLine(),
				Partial:  buf.GetPartial(),
				PartialLogMetadata: &PartialLogEntryMetadata{
					Last:    buf.PartialLogMetadata.GetLast(),
					Id:      buf.PartialLogMetadata.GetId(),
					Ordinal: buf.PartialLogMetadata.GetOrdinal(),
				},
			},
		}

		select {
		case lf.sendCh <- msg_proto:
			logrus.WithField("id", lf.info.ContainerID).Debugf("Message sent")
		default:
			// logrus.WithField("id", lf.info.ContainerID).Debugf("Message skipped")
		}

		buf.Reset()
	}
}

func RunService(lis net.Listener, d *driver) {
	var err error

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	clientfanout = fanout{mu: sync.RWMutex{}, clientchannel: make(map[uuid.UUID]chan *LogMessage)}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterIDockerLogDriverServer(grpcServer, d)
	reflection.Register(grpcServer)
	go d.fanoutmessages()
	grpcServer.Serve(lis)
}

func (d *driver) fanoutmessages() {
	for {
		msg := <-d.recvCh
		clientfanout.mu.RLock()

		for _, ch := range clientfanout.clientchannel {
			pcopy := &LogMessage{
				Service: msg.GetService(),
				Entry: &LogEntry{
					Source:   msg.GetEntry().GetSource(),
					TimeNano: msg.GetEntry().GetTimeNano(),
					Line:     msg.GetEntry().GetLine(),
					Partial:  msg.GetEntry().GetPartial(),
					PartialLogMetadata: &PartialLogEntryMetadata{
						Last:    msg.GetEntry().GetPartialLogMetadata().GetLast(),
						Id:      msg.GetEntry().GetPartialLogMetadata().GetId(),
						Ordinal: msg.GetEntry().GetPartialLogMetadata().GetOrdinal(),
					},
				},
			}
			ch <- pcopy
		}
		clientfanout.mu.RUnlock()
	}
}

func (d *driver) GetLogs(options *LogOptions, srv IDockerLogDriver_GetLogsServer) error {
	var msg *LogMessage
	id := uuid.New()

	clientfanout.mu.Lock()
	if _, ok := clientfanout.clientchannel[id]; ok {
		return status.Error(codes.Internal, "uuid clash")
	}
	mychannel := make(chan *LogMessage)
	clientfanout.clientchannel[id] = mychannel
	clientfanout.mu.Unlock()

LOOP:
	for {
		select {
		case msg = <-mychannel:
			srv.Send(msg)
		case <-d.stopCh:
			break LOOP
		}
	}

	clientfanout.mu.Lock()
	delete(clientfanout.clientchannel, id)
	clientfanout.mu.Unlock()

	return nil
}

func (d *driver) ListServices(context.Context, *emptypb.Empty) (*ServicesList, error) {
	slist := &ServicesList{}
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, lp := range d.logs {
		slist.Service = append(slist.Service, lp.info.ContainerName)
	}
	return slist, nil
}
