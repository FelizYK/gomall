package server

import (
	"context"
	"log"
	"net"

	"github.com/FelizYK/gomall/app/product/conf"
	"github.com/FelizYK/gomall/app/product/service"
	rpcproduct "github.com/FelizYK/gomall/rpc/product"
	"google.golang.org/grpc"
)

type productServiceServer struct {
	rpcproduct.UnimplementedProductServiceServer
}

func NewProductServiceServer() rpcproduct.ProductServiceServer {
	return &productServiceServer{}
}

func (s *productServiceServer) ListProducts(ctx context.Context, req *rpcproduct.ListProductsReq) (*rpcproduct.ListProductsResp, error) {
	return service.ListProducts(ctx, req)
}
func (s *productServiceServer) GetProduct(ctx context.Context, req *rpcproduct.GetProductReq) (*rpcproduct.GetProductResp, error) {
	return service.GetProduct(ctx, req)
}
func (s *productServiceServer) SearchProducts(ctx context.Context, req *rpcproduct.SearchProductsReq) (*rpcproduct.SearchProductsResp, error) {
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
	rpcproduct.RegisterProductServiceServer(s, &productServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func CloseServer() {
	if lis != nil {
		lis.Close()
	}
}
