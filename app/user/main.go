package main

import (
	"github.com/FelizYK/gomall/user/repository"
	"github.com/FelizYK/gomall/user/rpc"
)

func main() {
	repository.Init()

	rpc.InitEtcd()
	defer rpc.CloseEtcd()

	rpc.InitServer()
	defer rpc.CloseServer()
}
