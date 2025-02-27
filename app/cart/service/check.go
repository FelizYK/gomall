package service

import (
	"context"
	"errors"

	"github.com/FelizYK/gomall/app/cart/rpc"
	rpcproduct "github.com/FelizYK/gomall/rpc/product"
	rpcuser "github.com/FelizYK/gomall/rpc/user"
)

func CheckUser(ctx context.Context, userId uint32) error {
	user, err := rpc.UserClient.GetUser(ctx, &rpcuser.GetUserReq{
		Id: userId,
	})
	if err != nil {
		return err
	}
	if user.Email == "" {
		return errors.New("user not exist")
	}
	return nil
}

func CheckProduct(ctx context.Context, productId uint32) error {
	product, err := rpc.ProductClient.GetProduct(ctx, &rpcproduct.GetProductReq{
		Id: productId,
	})
	if err != nil {
		return err
	}
	if product.Product == nil || product.Product.Id == 0 {
		return errors.New("product not exist")
	}
	return nil
}
