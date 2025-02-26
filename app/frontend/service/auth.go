package service

import (
	"github.com/FelizYK/gomall/app/frontend/rpc"
	"github.com/FelizYK/gomall/app/frontend/rpc/auth"
	rpcuser "github.com/FelizYK/gomall/rpc/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context, req *auth.RegisterReq) (err error) {
	// call rpc
	resp, err := rpc.UserClient.Register(c, &rpcuser.RegisterReq{
		Email:           req.GetEmail(),
		Password:        req.GetPassword(),
		PasswordConfirm: req.GetPasswordConfirm(),
	})
	if err != nil {
		return err
	}
	// session
	session := sessions.Default(c)
	session.Set("user_id", resp.GetUserId())
	err = session.Save()
	return
}

func Login(c *gin.Context, req *auth.LoginReq) (err error) {
	// call rpc
	resp, err := rpc.UserClient.Login(c, &rpcuser.LoginReq{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return err
	}
	// session
	session := sessions.Default(c)
	session.Set("user_id", resp.GetUserId())
	err = session.Save()
	return
}

func Logout(c *gin.Context) (err error) {
	// session
	session := sessions.Default(c)
	session.Clear()
	err = session.Save()
	return
}
