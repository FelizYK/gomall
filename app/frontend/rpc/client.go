package rpc

import (
	"log"

	rpccart "github.com/FelizYK/gomall/rpc/cart"
	rpcproduct "github.com/FelizYK/gomall/rpc/product"
	rpcuser "github.com/FelizYK/gomall/rpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	UserClient    rpcuser.UserServiceClient
	ProductClient rpcproduct.ProductServiceClient
	CartClient    rpccart.CartServiceClient
)

func InitClient() {
	initUserClient()
	initProductClient()
	initCartClient()
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

func initCartClient() {
	cartAddr := DiscoverService("cart")
	conn, err := grpc.Dial(cartAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect cart service: %v", err)
	}
	CartClient = rpccart.NewCartServiceClient(conn)
}
