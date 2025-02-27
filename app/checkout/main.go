package main

import (
	"github.com/FelizYK/gomall/app/checkout/repository"
	"github.com/FelizYK/gomall/app/checkout/rpc"
	"github.com/FelizYK/gomall/app/checkout/rpc/server"
)

func main() {
	repository.Init()

	rpc.InitEtcd()
	defer rpc.CloseEtcd()

	server.InitServer()
	defer server.CloseServer()
}
