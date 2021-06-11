package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SomethingBot/gateway/api"
	"go.etcd.io/etcd/client/v3"
	"log"
	"sync"
	"time"
)

type EtcdConfig struct {
	EtcdUsername  string
	EtcdPassword  string
	EtcdEndpoints []string
}

type NodeInfo struct {
	sync.RWMutex
	NodeName   string
	MasterNode string
}

type Registry struct {
	Logger       *log.Logger
	EtcdConfig   EtcdConfig
	EtcdClient   *clientv3.Client
	EndpointName string
}

func New(logger *log.Logger, etcdConfig EtcdConfig) *Registry {
	return &Registry{Logger: logger, EtcdConfig: etcdConfig}
}

func (registry *Registry) Open() (err error) {
	registry.EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   registry.EtcdConfig.EtcdEndpoints,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return fmt.Errorf("could not open etcd client, %v", err)
	}
	return nil
}

func (registry *Registry) Close() (err error) {
	err = registry.EtcdClient.Close()
	if err != nil {
		return fmt.Errorf("could not close registry, %v", err)
	}
	return nil
}

//AddService into etcd, for apigateway
func (registry *Registry) AddService(service api.Service) (err error) {
	serviceJson, err := json.Marshal(service)
	//todo: handle error
	if err != nil {
		return err
	}

	_, err = registry.EtcdClient.Put(context.Background(), "services/"+service.Name, string(serviceJson))
	//todo: handle error
	if err != nil {
		return err
	}
	//
	//
	//resp, err := registry.EtcdClient.Get(context.Background(), "services/"+service.Name+"/*")
	//if err != nil {
	//	return nil
	//}
	//registry.Logger.Printf("GET services/"+service.Name+": +%v", resp)

	return nil
}
