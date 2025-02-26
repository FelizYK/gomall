package service

import (
	"context"

	"github.com/FelizYK/gomall/app/user/repository"
	rpcuser "github.com/FelizYK/gomall/rpc/user"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx context.Context, req *rpcuser.LoginReq) (resp *rpcuser.LoginResp, err error) {
	// get user by email
	userRow, err := repository.GetByEmail(ctx, req.Email)
	if err != nil {
		return
	}
	// check password
	if bcrypt.CompareHashAndPassword([]byte(userRow.PasswordHashed), []byte(req.Password)); err != nil {
		return
	}
	return &rpcuser.LoginResp{UserId: int32(userRow.ID)}, nil
}
