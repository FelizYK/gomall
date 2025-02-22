package main

import (
	"github.com/FelizYK/gomall/user/repository"
	"github.com/FelizYK/gomall/user/rpc"
)

func main() {
	repository.Init()

	go rpc.InitServer()
	defer rpc.CloseServer()

	rpc.InitEtcd()
	defer rpc.CloseEtcd()
}
