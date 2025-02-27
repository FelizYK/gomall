package service

import (
	"github.com/FelizYK/gomall/app/frontend/rpc"
	"github.com/FelizYK/gomall/app/frontend/rpc/product"
	rpcproduct "github.com/FelizYK/gomall/rpc/product"
	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context, req *product.GetProductReq) (map[string]any, error) {
	// call rpc
	resp, err := rpc.ProductClient.GetProduct(c, &rpcproduct.GetProductReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	// assemble response
	return gin.H{
		"id":          resp.Product.Id,
		"name":        resp.Product.Name,
		"description": resp.Product.Description,
		"picture":     resp.Product.Picture,
		"price":       resp.Product.Price,
	}, nil
}

func ListProductsByCategory(c *gin.Context, req *product.ListProductsReq) (map[string]any, error) {
	// call rpc
	resp, err := rpc.ProductClient.ListProducts(c, &rpcproduct.ListProductsReq{
		CategoryName: req.Category,
	})
	if err != nil {
		return nil, err
	}
	// assemble response
	var products []map[string]any
	for _, p := range resp.Products {
		products = append(products, gin.H{
			"id":          p.Id,
			"name":        p.Name,
			"description": p.Description,
			"picture":     p.Picture,
			"price":       p.Price,
		})
	}
	return gin.H{
		"products": products,
	}, nil
}

func SearchProducts(c *gin.Context, req *product.SearchProductsReq) (map[string]any, error) {
	// call rpc
	resp, err := rpc.ProductClient.SearchProducts(c, &rpcproduct.SearchProductsReq{
		Query: req.Query,
	})
	if err != nil {
		return nil, err
	}
	// assemble response
	var products []map[string]any
	for _, p := range resp.Products {
		products = append(products, gin.H{
			"id":          p.Id,
			"name":        p.Name,
			"description": p.Description,
			"picture":     p.Picture,
			"price":       p.Price,
		})
	}
	return gin.H{
		"query":    req.Query,
		"products": products,
	}, nil
}
