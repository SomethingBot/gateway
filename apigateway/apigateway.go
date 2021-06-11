package apigateway

import (
	"fmt"
	"github.com/SomethingBot/gateway/api"
	"github.com/SomethingBot/gateway/registry"
	"log"
	"net/http"
	"sync"
	"time"
)

type ApiGateway struct {
	logger     *log.Logger
	registry   *registry.Registry
	httpServer *http.Server
	serveMux   *http.ServeMux
	Address    string
	waitGroup  *sync.WaitGroup
}

func New(logger *log.Logger, etcdConfig registry.EtcdConfig, address string) (apiGateway ApiGateway) {
	apiGateway.logger = logger
	apiGateway.registry = registry.New(logger, etcdConfig)
	apiGateway.Address = address
	apiGateway.waitGroup = &sync.WaitGroup{}
	return
}

func (apiGateway *ApiGateway) Open() (err error) {
	err = apiGateway.registry.Open()
	if err != nil {
		return
	}

	err = apiGateway.registry.AddService(api.Service{
		Name: "gateway",
		Commands: []api.Command{{
			Name:   "statistics",
			Prefix: false,
		}},
	})
	if err != nil {
		return err
	}

	apiGateway.serveMux = &http.ServeMux{}
	apiGateway.httpServer = &http.Server{
		Addr:              apiGateway.Address,
		Handler:           apiGateway.serveMux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	apiGateway.serveMux.HandleFunc("/", apiGateway.ServeHTTP)

	errChan := make(chan error)

	apiGateway.waitGroup.Add(1)
	go apiGateway.run(errChan)

	select {
	case err = <-errChan:
		return err
	case <-time.After(1 * time.Second):
		return nil
	}
}

func (apiGateway *ApiGateway) Close() (err error) {
	err = apiGateway.httpServer.Close()
	if err != nil {
		return err
	}

	apiGateway.waitGroup.Wait()

	err = apiGateway.registry.Close()
	if err != nil {
		return
	}

	return nil
}

func (apiGateway *ApiGateway) run(errChan chan error) {
	err := apiGateway.httpServer.ListenAndServe()
	if err != nil {
		errChan <- err
	}
	apiGateway.waitGroup.Done()
}

func (apiGateway *ApiGateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello\n")
	if err != nil {
		apiGateway.logger.Printf("Could not write http response err: %v", err)
	}
}
