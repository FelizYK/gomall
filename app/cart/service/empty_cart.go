package service

import (
	"context"

	"github.com/FelizYK/gomall/app/cart/repository"
	rpccart "github.com/FelizYK/gomall/rpc/cart"
)

func EmptyCart(ctx context.Context, req *rpccart.EmptyCartReq) (err error) {
	// check user exists
	err = CheckUser(ctx, req.GetUserId())
	// empty cart by user_id
	err = repository.EmptyCart(ctx, uint(req.GetUserId()))
	if err != nil {
		return
	}
	return
}
