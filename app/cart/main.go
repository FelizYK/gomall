package main

import (
	"github.com/FelizYK/gomall/app/cart/repository"
	"github.com/FelizYK/gomall/app/cart/rpc"
	"github.com/FelizYK/gomall/app/cart/rpc/server"
)

func main() {
	repository.Init()

	rpc.InitEtcd()
	defer rpc.CloseEtcd()

	server.InitServer()
	defer server.CloseServer()
}
