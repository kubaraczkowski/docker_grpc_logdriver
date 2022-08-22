package driver

import (
	"context"
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/docker/docker/api/types/plugins/logdriver"
	"github.com/docker/docker/daemon/logger"
	"github.com/stretchr/testify/assert"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestSendMsg(t *testing.T) {
	lis = bufconn.Listen(bufSize)

	d := NewDriver()

	lf := &logPair{nil, logger.Info{}, d.recvCh, d.stopCh, make(chan logdriver.LogEntry)}
	go consumeLog(lf)

	go RunService(lis, d)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	received := make(chan struct{})

	// client
	go func() {
		client := NewIDockerLogDriverClient(conn)
		stream, err := client.GetLogs(context.Background(), &LogOptions{})
		if err != nil {
			t.Fail()
		}
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fail()
			}
			fmt.Print(msg.String())
			if assert.Equal(t, "test", msg.Entry.GetSource()) {
				close(received)
				break
			}
		}

	}()

CLLIENT_LOOP:
	for {
		select {
		case lf.msgCh <- logdriver.LogEntry{Source: "test"}:
		case <-received:
			break CLLIENT_LOOP
		}
	}

	close(d.stopCh)
}
