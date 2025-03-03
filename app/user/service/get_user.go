package service

import (
	"context"

	"github.com/FelizYK/gomall/app/user/repository"
	rpcuser "github.com/FelizYK/gomall/rpc/user"
)

func GetUser(ctx context.Context, req *rpcuser.GetUserReq) (resp *rpcuser.GetUserResp, err error) {
	// get user by id
	q := repository.NewUserQuery(ctx)
	user, err := q.GetById(req.Id)
	if err != nil {
		return
	}
	// assemble response
	return &rpcuser.GetUserResp{Email: user.Email}, nil
}
