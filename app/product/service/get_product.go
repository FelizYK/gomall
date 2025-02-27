package service

import (
	"context"
	"errors"

	"github.com/FelizYK/gomall/app/product/repository"
	rpcproduct "github.com/FelizYK/gomall/rpc/product"
)

func GetProduct(ctx context.Context, req *rpcproduct.GetProductReq) (resp *rpcproduct.GetProductResp, err error) {
	// check if id exists
	if req.Id == 0 {
		return nil, errors.New("product id is required")
	}
	// get product by id
	product, err := repository.GetProductById(ctx, req.Id)
	if err != nil {
		return
	}
	// assemble product
	return &rpcproduct.GetProductResp{
		Product: &rpcproduct.Product{
			Id:          uint32(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Picture:     product.Picture,
			Price:       product.Price,
		},
	}, nil
}
