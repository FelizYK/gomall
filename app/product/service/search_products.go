package service

import (
	"context"

	"github.com/FelizYK/gomall/app/product/repository"
	productrpc "github.com/FelizYK/gomall/rpc/product"
)

func SearchProducts(ctx context.Context, req *productrpc.SearchProductsReq) (resp *productrpc.SearchProductsResp, err error) {
	// search products
	products, err := repository.SearchProducts(ctx, req.Query)
	if err != nil {
		return
	}
	// assemble products
	resp = &productrpc.SearchProductsResp{}
	for _, p := range products {
		resp.Products = append(resp.Products, &productrpc.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		})
	}
	return
}
