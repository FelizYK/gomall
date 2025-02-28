package server

import (
	"context"
	"log"
	"net"

	"github.com/FelizYK/gomall/app/order/conf"
	"github.com/FelizYK/gomall/app/order/service"
	rpcorder "github.com/FelizYK/gomall/rpc/order"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type orderServiceServer struct {
	rpcorder.UnimplementedOrderServiceServer
}

func NewOrderServiceServer() rpcorder.OrderServiceServer {
	return &orderServiceServer{}
}

func (s *orderServiceServer) AddOrder(ctx context.Context, req *rpcorder.AddOrderReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, service.AddOrder(ctx, req)
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
	rpcorder.RegisterOrderServiceServer(s, &orderServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func CloseServer() {
	if lis != nil {
		lis.Close()
	}
}
