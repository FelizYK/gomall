package service

import (
	"context"
	"errors"

	"github.com/FelizYK/gomall/app/user/repository"
	rpcuser "github.com/FelizYK/gomall/rpc/user"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx context.Context, req *rpcuser.RegisterReq) (resp *rpcuser.RegisterResp, err error) {
	// check password
	if req.Password != req.PasswordConfirm {
		err = errors.New("different Password and PasswordConfirm")
		return
	}
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	// create user
	newUser := &repository.User{
		Email:          req.Email,
		PasswordHashed: string(hashedPassword),
	}
	q := repository.NewUserQuery(ctx)
	if err = q.Create(newUser); err != nil {
		return
	}
	// assemble response
	return &rpcuser.RegisterResp{UserId: uint32(newUser.ID)}, nil
}
