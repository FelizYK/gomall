package server

import (
	"context"
	"log"
	"net"

	"github.com/FelizYK/gomall/app/user/conf"
	"github.com/FelizYK/gomall/app/user/service"
	rpcuser "github.com/FelizYK/gomall/rpc/user"
	"google.golang.org/grpc"
)

type userServiceServer struct {
	rpcuser.UnimplementedUserServiceServer
}

func NewUserServiceServer() rpcuser.UserServiceServer {
	return &userServiceServer{}
}

func (s *userServiceServer) Register(ctx context.Context, req *rpcuser.RegisterReq) (*rpcuser.RegisterResp, error) {
	return service.Register(ctx, req)
}
func (s *userServiceServer) Login(ctx context.Context, req *rpcuser.LoginReq) (*rpcuser.LoginResp, error) {
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
	rpcuser.RegisterUserServiceServer(s, &userServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func CloseServer() {
	if lis != nil {
		lis.Close()
	}
}
