package service

import (
	"context"

	"github.com/FelizYK/gomall/app/product/repository"
	rpcproduct "github.com/FelizYK/gomall/rpc/product"
)

func SearchProducts(ctx context.Context, req *rpcproduct.SearchProductsReq) (resp *rpcproduct.SearchProductsResp, err error) {
	// search products
	q := repository.NewProductQuery(ctx)
	products, err := q.SearchProducts(req.Query)
	if err != nil {
		return
	}
	// assemble products
	resp = &rpcproduct.SearchProductsResp{}
	for _, p := range products {
		resp.Products = append(resp.Products, &rpcproduct.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		})
	}
	return
}
