package rpc

import (
	"log"

	rpccart "github.com/FelizYK/gomall/rpc/cart"
	rpcproduct "github.com/FelizYK/gomall/rpc/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	CartClient    rpccart.CartServiceClient
	ProductClient rpcproduct.ProductServiceClient
)

func initClient() {
	initCartClient()
	initProductClient()
}

func initCartClient() {
	userAddr := discoverService("cart")
	conn, err := grpc.Dial(userAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect user service: %v", err)
	}
	CartClient = rpccart.NewCartServiceClient(conn)
}

func initProductClient() {
	userAddr := discoverService("product")
	conn, err := grpc.Dial(userAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect user service: %v", err)
	}
	ProductClient = rpcproduct.NewProductServiceClient(conn)
}
