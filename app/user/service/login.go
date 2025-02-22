package service

import (
	"context"

	userrpc "github.com/FelizYK/gomall/rpc/user"
	"github.com/FelizYK/gomall/user/repository"
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
