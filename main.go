package main

import (
	"github.com/SomethingBot/gateway/registry"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	GitCommit string
	GitTag    string
	Mode      string
)

func parseEnv() (config registry.Config) {
	config.EtcdUsername = os.Getenv("GATEWAY_ETCD_USERNAME")
	config.EtcdPassword = os.Getenv("GATEWAY_ETCD_PASSWORD")
	config.EtcdEndpoints = strings.Split(os.Getenv("GATEWAY_ETCD_ENDPOINTS"), ",")
	return
}

func main() {
	logger := log.New(os.Stdout, "", log.LUTC|log.Ldate|log.Ltime)
	logger.Printf("Starting gateway, commit:%v, tag:%v, Mode:%v", GitCommit, GitTag, Mode)

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	signalFromSystem := <-osSignal

	logger.Printf("Stopping gateway, reason: signal-%v", signalFromSystem.String())
}
