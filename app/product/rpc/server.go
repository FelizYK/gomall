package rpc

import (
	"context"
	"log"
	"net"

	"github.com/FelizYK/gomall/app/product/conf"
	"github.com/FelizYK/gomall/app/product/service"
	productrpc "github.com/FelizYK/gomall/rpc/product"
	"google.golang.org/grpc"
)

type productServiceServer struct {
	productrpc.UnimplementedProductServiceServer
}

func NewProductServiceServer() productrpc.ProductServiceServer {
	return &productServiceServer{}
}

func (s *productServiceServer) ListProducts(ctx context.Context, req *productrpc.ListProductsReq) (*productrpc.ListProductsResp, error) {
	return service.ListProducts(ctx, req)
}
func (s *productServiceServer) GetProduct(ctx context.Context, req *productrpc.GetProductReq) (*productrpc.GetProductResp, error) {
	return service.GetProduct(ctx, req)
}
func (s *productServiceServer) SearchProducts(ctx context.Context, req *productrpc.SearchProductsReq) (*productrpc.SearchProductsResp, error) {
	return service.SearchProducts(ctx, req)
}

var (
	lis net.Listener
	err error
)

func InitServer() {
	lis, err = net.Listen("tcp", conf.GetConf().Service.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	productrpc.RegisterProductServiceServer(s, &productServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func CloseServer() {
	if lis != nil {
		lis.Close()
	}
}
