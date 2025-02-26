package service

import (
	"context"

	"github.com/FelizYK/gomall/app/product/repository"
	rpcproduct "github.com/FelizYK/gomall/rpc/product"
)

func ListProducts(ctx context.Context, req *rpcproduct.ListProductsReq) (resp *rpcproduct.ListProductsResp, err error) {
	if req.GetCategoryName() == "" {
		return getProducts(ctx)
	} else {
		return getProductsByCategory(ctx, req)
	}
}

func getProducts(ctx context.Context) (resp *rpcproduct.ListProductsResp, err error) {
	// get category by name
	products, err := repository.GetProducts(ctx)
	if err != nil {
		return
	}
	// assemble products
	resp = &rpcproduct.ListProductsResp{}
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

func getProductsByCategory(ctx context.Context, req *rpcproduct.ListProductsReq) (resp *rpcproduct.ListProductsResp, err error) {
	// get category by name
	category, err := repository.GetCategoryByName(ctx, req.CategoryName)
	if err != nil {
		return
	}
	// assemble products
	resp = &rpcproduct.ListProductsResp{}
	for _, p := range category.Products {
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
