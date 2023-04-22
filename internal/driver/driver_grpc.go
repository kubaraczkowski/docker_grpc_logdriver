package driver

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/containerd/fifo"
	"github.com/docker/docker/api/types/plugins/logdriver"
	"github.com/docker/docker/daemon/logger"
	protoio "github.com/gogo/protobuf/io"
	"github.com/google/uuid"
	"github.com/natefinch/lumberjack"
	"github.com/pkg/errors"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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
	stream    io.ReadCloser
	info      logger.Info
	sendCh    chan *LogMessage
	stopCh    chan struct{}
	msgCh     chan logdriver.LogEntry
	msgFileCh chan logdriver.LogEntry
}

type fanout struct {
	mu            sync.RWMutex
	clientchannel map[uuid.UUID](chan *LogMessage)
}

var clientfanout fanout

func NewDriver() (*driver, error) {
	return &driver{
		logs:   make(map[string]*logPair),
		idx:    make(map[string]*logPair),
		recvCh: make(chan *LogMessage),
		stopCh: make(chan struct{}),
	}, nil
}

func (d *driver) StartLogging(file string, logCtx logger.Info) error {
	log.Printf("StarLogging container %s", logCtx.ContainerID)
	d.mu.Lock()
	if _, exists := d.logs[file]; exists {
		d.mu.Unlock()
		return fmt.Errorf("logger for %q already exists", file)
	}
	d.mu.Unlock()

	f, err := fifo.OpenFifo(context.Background(), file, syscall.O_RDONLY, 0700)
	if err != nil {
		return errors.Wrapf(err, "error opening logger fifo: %q", file)
	}

	d.mu.Lock()
	lf := &logPair{
		stream:    f,
		info:      logCtx,
		sendCh:    d.recvCh,
		stopCh:    d.stopCh,
		msgCh:     make(chan logdriver.LogEntry),
		msgFileCh: make(chan logdriver.LogEntry),
	}
	d.logs[file] = lf
	d.idx[logCtx.ContainerID] = lf
	d.mu.Unlock()

	go readLog(lf)
	go consumeLog(lf)
	go writeToFile(lf)
	return nil
}

func (d *driver) StopLogging(file string) error {
	log.Print("StopLogging")
	d.mu.Lock()
	lf, ok := d.logs[file]
	if ok {
		lf.stream.Close()
		delete(d.logs, file)
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
			if err != nil {
				if err == io.EOF {
				} else {
					log.Print(err.Error())
				}
				lf.stream.Close()
				return
			}
			dec = protoio.NewUint32DelimitedReader(lf.stream, binary.BigEndian, 1e6)
			continue
		}
		lf.msgCh <- buf
		lf.msgFileCh <- buf
	}
}

func consumeLog(lf *logPair) {

	for {

		buf, ok := <-lf.msgCh
		if !ok {
			return
		}

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
		default:
			// logrus.WithField("id", lf.info.ContainerID).Debugf("Message skipped")
		}

		buf.Reset()
	}
}

var logFileBaseDir = "/log"

func writeToFile(lf *logPair) {
	// setup a rotating logger based on the container name
	name := lf.info.ContainerName
	logger := &lumberjack.Logger{
		Filename:   filepath.Join(logFileBaseDir, fmt.Sprintf("%s.log", name)),
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}
	defer logger.Close()

	var msg logdriver.LogEntry
	var ok bool
	for {
		msg, ok = <-lf.msgFileCh
		if !ok {
			return
		}
		logger.Write(msg.Line)
		logger.Write([]byte("\n"))
	}

}

func RunService(lis net.Listener, d *driver) {
	var err error

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	clientfanout = fanout{mu: sync.RWMutex{}, clientchannel: make(map[uuid.UUID]chan *LogMessage)}
	defer close(d.stopCh)
	defer close(d.recvCh)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterIDockerLogDriverServer(grpcServer, d)
	reflection.Register(grpcServer)
	go d.fanoutmessages()
	grpcServer.Serve(lis)

}

func (d *driver) fanoutmessages() {
	for {
		msg, ok := <-d.recvCh
		if !ok {
			// channel closed
			return
		}
		clientfanout.mu.RLock()

		for _, ch := range clientfanout.clientchannel {
			pcopy := proto.Clone(msg)
			ch <- pcopy.(*LogMessage)
		}
		clientfanout.mu.RUnlock()
	}
}

func (d *driver) GetLogs(options *LogOptions, srv IDockerLogDriver_GetLogsServer) error {
	var msg *LogMessage
	id := uuid.New()
	log.Printf("New client: %v", id)

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
