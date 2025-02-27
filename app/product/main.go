package main

import (
	"github.com/FelizYK/gomall/app/product/repository"
	"github.com/FelizYK/gomall/app/product/rpc"
	"github.com/FelizYK/gomall/app/product/rpc/server"
)

func main() {
	repository.Init()

	rpc.InitEtcd()
	defer rpc.CloseEtcd()

	server.InitServer()
	defer server.CloseServer()
}
