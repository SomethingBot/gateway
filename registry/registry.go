package registry

import (
	"go.etcd.io/etcd/client/v3"
)

type Config struct {
	EtcdUsername  string
	EtcdPassword  string
	EtcdEndpoints []string
}

type Registry struct {
	EtcdConfig Config
	EtcdClient clientv3.Client
}

func New(config Config) (Registry, error) {
	return Registry{EtcdConfig: config}, nil
}
