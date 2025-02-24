package main

import (
	"github.com/FelizYK/gomall/app/user/repository"
	"github.com/FelizYK/gomall/app/user/rpc"
)

func main() {
	repository.Init()

	rpc.InitEtcd()
	defer rpc.CloseEtcd()

	rpc.InitServer()
	defer rpc.CloseServer()
}
