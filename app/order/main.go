package main

import (
	"github.com/FelizYK/gomall/app/order/repository"
	"github.com/FelizYK/gomall/app/order/rpc"
	"github.com/FelizYK/gomall/app/order/rpc/server"
)

func main() {
	repository.Init()

	rpc.InitEtcd()
	defer rpc.CloseEtcd()

	server.InitServer()
	defer server.CloseServer()
}
