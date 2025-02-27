package main

import (
	"github.com/FelizYK/gomall/app/user/repository"
	"github.com/FelizYK/gomall/app/user/rpc"
	"github.com/FelizYK/gomall/app/user/rpc/server"
)

func main() {
	repository.Init()

	rpc.InitEtcd()
	defer rpc.CloseEtcd()

	server.InitServer()
	defer server.CloseServer()
}
