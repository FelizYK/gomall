package service

import (
	"github.com/FelizYK/gomall/app/frontend/rpc"
	"github.com/FelizYK/gomall/app/frontend/rpc/auth"
	userrpc "github.com/FelizYK/gomall/rpc/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Register(c *gin.Context, req *auth.RegisterReq) (*emptypb.Empty, error) {
	// call rpc
	resp, err := rpc.UserClient.Register(c, &userrpc.RegisterReq{
		Email:           req.GetEmail(),
		Password:        req.GetPassword(),
		PasswordConfirm: req.GetPasswordConfirm(),
	})
	if err != nil {
		return nil, err
	}
	// session
	session := sessions.Default(c)
	session.Set("user_id", resp.GetUserId())
	if err = session.Save(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func Login(c *gin.Context, req *auth.LoginReq) (*emptypb.Empty, error) {
	// call rpc
	resp, err := rpc.UserClient.Login(c, &userrpc.LoginReq{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	// session
	session := sessions.Default(c)
	session.Set("user_id", resp.GetUserId())
	if err = session.Save(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func Logout(c *gin.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	// session
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
