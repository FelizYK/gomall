package rpc

import (
	"context"
	"log"
	"time"

	"github.com/FelizYK/gomall/app/frontend/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	cli *clientv3.Client
)

func InitEtcd() {
	var err error
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{conf.GetConf().Etcd.Endpoints},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to connect etcd: %v", err)
	}
}

func CloseEtcd() {
	if cli != nil {
		cli.Close()
	}
}

func DiscoverService(serviceName string) (serviceAddr string) {
	resp, err := cli.Get(context.TODO(), serviceName, clientv3.WithPrefix())
	if err != nil {
		log.Fatalf("Failed to get service address: %v", err)
	}
	if len(resp.Kvs) == 0 {
		log.Fatalf("No service found: %s", serviceName)
	}
	return string(resp.Kvs[0].Value)
}
