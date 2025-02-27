package service

import (
	"context"

	"github.com/FelizYK/gomall/app/cart/repository"
	rpccart "github.com/FelizYK/gomall/rpc/cart"
)

func GetCart(ctx context.Context, req *rpccart.GetCartReq) (resp *rpccart.GetCartResp, err error) {
	// check user exists
	err = CheckUser(ctx, req.GetUserId())
	// get cart_items by user_id
	items, err := repository.GetCartByUserId(ctx, uint(req.GetUserId()))
	if err != nil {
		return
	}
	// assemble response
	resp = &rpccart.GetCartResp{}
	for _, item := range items {
		resp.Items = append(resp.Items, &rpccart.CartItem{
			ProductId: uint32(item.ProductId),
			Quantity:  int32(item.Quantity),
		})
	}
	return
}
