package rpc

import (
	"context"
	"log"
	"time"

	"github.com/FelizYK/gomall/app/user/conf"
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

	registerService()
}

func CloseEtcd() {
	if cli != nil {
		cli.Close()
	}
}

func registerService() {
	// grant lease
	ctx := context.Background()
	lease, err := cli.Grant(ctx, 30)
	if err != nil {
		log.Fatalf("Failed to create lease: %v", err)
	}
	// register service
	serviceName := conf.GetConf().Service.Name
	serviceAddress := conf.GetConf().Service.Address
	_, err = cli.Put(ctx, serviceName, serviceAddress, clientv3.WithLease(lease.ID))
	if err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}
	// keep alive lease
	keepAliveChan, err := cli.KeepAlive(ctx, lease.ID)
	if err != nil {
		log.Fatalf("Failed to keep alive lease: %v", err)
	}
	go func() {
		for {
			<-keepAliveChan
		}
	}()
}
