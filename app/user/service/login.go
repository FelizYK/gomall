package service

import (
	"context"

	"github.com/FelizYK/gomall/app/user/repository"
	rpcuser "github.com/FelizYK/gomall/rpc/user"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx context.Context, req *rpcuser.LoginReq) (resp *rpcuser.LoginResp, err error) {
	// get user by email
	user, err := repository.GetByEmail(ctx, req.Email)
	if err != nil {
		return
	}
	// check password
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHashed), []byte(req.Password)); err != nil {
		return
	}
	// assemble response
	return &rpcuser.LoginResp{UserId: uint32(user.ID)}, nil
}
