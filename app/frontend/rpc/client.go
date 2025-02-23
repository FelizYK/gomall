package rpc

import (
	"log"

	userrpc "github.com/FelizYK/gomall/rpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	UserClient userrpc.UserServiceClient
)

func InitClient() {
	initUserClient()
}

func initUserClient() {
	userAddr := DiscoverService("user")
	conn, err := grpc.Dial(userAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect user service: %v", err)
	}
	UserClient = userrpc.NewUserServiceClient(conn)
}
