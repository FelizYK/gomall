package rpc

import (
	"log"

	rpcproduct "github.com/FelizYK/gomall/rpc/product"
	rpcuser "github.com/FelizYK/gomall/rpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	UserClient    rpcuser.UserServiceClient
	ProductClient rpcproduct.ProductServiceClient
)

func InitClient() {
	initUserClient()
	initProductClient()
}

func initUserClient() {
	userAddr := DiscoverService("user")
	conn, err := grpc.Dial(userAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect user service: %v", err)
	}
	UserClient = rpcuser.NewUserServiceClient(conn)
}

func initProductClient() {
	productAddr := DiscoverService("product")
	conn, err := grpc.Dial(productAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect product service: %v", err)
	}
	ProductClient = rpcproduct.NewProductServiceClient(conn)
}
