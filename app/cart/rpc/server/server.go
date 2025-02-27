package server

import (
	"context"
	"log"
	"net"

	"github.com/FelizYK/gomall/app/cart/conf"
	"github.com/FelizYK/gomall/app/cart/service"
	rpccart "github.com/FelizYK/gomall/rpc/cart"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type cartServiceServer struct {
	rpccart.UnimplementedCartServiceServer
}

func NewCartServiceServer() rpccart.CartServiceServer {
	return &cartServiceServer{}
}

func (s *cartServiceServer) GetCart(ctx context.Context, req *rpccart.GetCartReq) (*rpccart.GetCartResp, error) {
	return service.GetCart(ctx, req)
}
func (s *cartServiceServer) AddCart(ctx context.Context, req *rpccart.AddCartReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, service.AddCart(ctx, req)
}
func (s *cartServiceServer) EmptyCart(ctx context.Context, req *rpccart.EmptyCartReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, service.EmptyCart(ctx, req)
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
	rpccart.RegisterCartServiceServer(s, &cartServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func CloseServer() {
	if lis != nil {
		lis.Close()
	}
}
