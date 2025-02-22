package rpc

import (
	"context"
	"log"
	"net"

	userrpc "github.com/FelizYK/gomall/rpc/user"
	"github.com/FelizYK/gomall/user/conf"
	"github.com/FelizYK/gomall/user/service"
	"google.golang.org/grpc"
)

type userServiceServer struct {
	userrpc.UnimplementedUserServiceServer
}

func NewUserServiceServer() userrpc.UserServiceServer {
	return &userServiceServer{}
}

func (s *userServiceServer) Register(ctx context.Context, req *userrpc.RegisterReq) (*userrpc.RegisterResp, error) {
	return service.Register(ctx, req)
}
func (s *userServiceServer) Login(ctx context.Context, req *userrpc.LoginReq) (*userrpc.LoginResp, error) {
	return service.Login(ctx, req)
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
	userrpc.RegisterUserServiceServer(s, &userServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func CloseServer() {
	if lis != nil {
		lis.Close()
	}
}
