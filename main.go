package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	_ "embed"

	"github.com/docker/go-plugins-helpers/sdk"
	"github.com/kubaraczkowski/docker_grpc_logdriver/internal/driver"
	"github.com/sirupsen/logrus"
)

var logLevels = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

var (
	port = flag.Int("port", 50055, "The server port")
)

const socketAddress = "/run/docker/plugins/grpc.sock"

//go:embed version.txt
var version string

func main() {
	flag.Parse()

	log.Printf("GRPC LogDriver version %s", version)

	levelVal := os.Getenv("LOG_LEVEL")
	if levelVal == "" {
		levelVal = "info"
	}
	if level, exists := logLevels[levelVal]; exists {
		logrus.SetLevel(level)
	} else {
		fmt.Fprintln(os.Stderr, "invalid log level: ", levelVal)
		os.Exit(1)
	}

	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)                     //default
	log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout

	h := sdk.NewHandler(`{"Implements": ["LoggingDriver"]}`)
	d := driver.NewDriver()
	driver.Handlers(&h, d)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	go driver.RunService(lis, d)

	if err := h.ServeUnix(socketAddress, 0); err != nil {
		panic(err)
	}

}
