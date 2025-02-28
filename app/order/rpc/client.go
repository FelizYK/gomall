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
	cartAddr := discoverService("cart")
	conn, err := grpc.Dial(cartAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect cart service: %v", err)
	}
	CartClient = rpccart.NewCartServiceClient(conn)
}

func initProductClient() {
	productAddr := discoverService("product")
	conn, err := grpc.Dial(productAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect product service: %v", err)
	}
	ProductClient = rpcproduct.NewProductServiceClient(conn)
}
