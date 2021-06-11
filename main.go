package main

import (
	"github.com/SomethingBot/gateway/apigateway"
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

func parseEtcdConfig() (config registry.EtcdConfig) {
	config.EtcdUsername = os.Getenv("GATEWAY_ETCD_USERNAME")
	config.EtcdPassword = os.Getenv("GATEWAY_ETCD_PASSWORD")
	config.EtcdEndpoints = strings.Split(os.Getenv("GATEWAY_ETCD_ENDPOINTS"), ",")
	return
}

func parseApiGatewayConfig() string {
	return os.Getenv("GATEWAY_ADDRESS")
}

func main() {
	logger := log.New(os.Stdout, "", log.LUTC|log.Ldate|log.Ltime)
	logger.Printf("Starting gateway, commit:%v, tag:%v, Mode:%v", GitCommit, GitTag, Mode)

	apiGateway := apigateway.New(logger, parseEtcdConfig(), parseApiGatewayConfig())
	err := apiGateway.Open()
	if err != nil {
		logger.Printf("Could not start, reason: could not open ApiGateway, err: %v", err)
	}

	logger.Printf("Started Gateway on %v", apiGateway.Address)

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	if signalFromSystem := <-osSignal; signalFromSystem != nil {
		logger.Printf("Stopping gateway, reason: signal-%v", signalFromSystem.String())
		return
	}

	err = apiGateway.Close()
	if err != nil {
		logger.Printf("Error stopping gateway, %v", err)
		return
	}

	logger.Printf("Stopped gateway")
}
