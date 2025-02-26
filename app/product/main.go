package main

import (
	"github.com/FelizYK/gomall/app/product/repository"
	"github.com/FelizYK/gomall/app/product/rpc"
)

func main() {
	repository.Init()

	rpc.InitEtcd()
	defer rpc.CloseEtcd()

	rpc.InitServer()
	defer rpc.CloseServer()
}
