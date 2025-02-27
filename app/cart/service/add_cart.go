package service

import (
	"context"

	"github.com/FelizYK/gomall/app/cart/repository"
	rpccart "github.com/FelizYK/gomall/rpc/cart"
)

func AddCart(ctx context.Context, req *rpccart.AddCartReq) (err error) {
	// check user exists
	err = CheckUser(ctx, req.UserId)
	if err != nil {
		return
	}
	// check product exists
	err = CheckProduct(ctx, req.Item.ProductId)
	if err != nil {
		return
	}
	// add cart_item
	err = repository.AddCart(ctx, uint(req.UserId), uint(req.Item.ProductId), int(req.Item.Quantity))
	if err != nil {
		return
	}
	return
}
