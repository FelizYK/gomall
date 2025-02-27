package server

import (
	"context"
	"log"
	"net"

	"github.com/FelizYK/gomall/app/checkout/conf"
	"github.com/FelizYK/gomall/app/checkout/service"
	rpccheckout "github.com/FelizYK/gomall/rpc/checkout"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type checkoutServiceServer struct {
	rpccheckout.UnimplementedCheckoutServiceServer
}

func NewCheckoutServiceServer() rpccheckout.CheckoutServiceServer {
	return &checkoutServiceServer{}
}

func (s *checkoutServiceServer) Checkout(ctx context.Context, req *rpccheckout.CheckoutReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, service.Checkout(ctx, req)
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
	rpccheckout.RegisterCheckoutServiceServer(s, &checkoutServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func CloseServer() {
	if lis != nil {
		lis.Close()
	}
}
