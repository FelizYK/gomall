package service

import (
	"context"

	"github.com/FelizYK/gomall/app/user/repository"
	userrpc "github.com/FelizYK/gomall/rpc/user"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx context.Context, req *userrpc.LoginReq) (resp *userrpc.LoginResp, err error) {
	// get user by email
	userRow, err := repository.GetByEmail(ctx, req.Email)
	if err != nil {
		return
	}
	// check password
	if bcrypt.CompareHashAndPassword([]byte(userRow.PasswordHashed), []byte(req.Password)); err != nil {
		return
	}
	return &userrpc.LoginResp{UserId: int32(userRow.ID)}, nil
}
