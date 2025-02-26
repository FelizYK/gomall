package service

import (
	"context"

	"github.com/FelizYK/gomall/app/product/repository"
	productrpc "github.com/FelizYK/gomall/rpc/product"
)

func ListProducts(ctx context.Context, req *productrpc.ListProductsReq) (resp *productrpc.ListProductsResp, err error) {
	// get category by name
	category, err := repository.GetCategoryByName(ctx, req.CategoryName)
	if err != nil {
		return
	}
	// assemble products
	resp = &productrpc.ListProductsResp{}
	for _, p := range category.Products {
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
