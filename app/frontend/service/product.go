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
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	// assemble response
	return gin.H{
		"id":          resp.GetProduct().GetId(),
		"name":        resp.GetProduct().GetName(),
		"description": resp.GetProduct().GetDescription(),
		"picture":     resp.GetProduct().GetPicture(),
		"price":       resp.GetProduct().GetPrice(),
	}, nil
}

func ListProductsByCategory(c *gin.Context, req *product.ListProductsReq) (map[string]any, error) {
	// call rpc
	resp, err := rpc.ProductClient.ListProducts(c, &rpcproduct.ListProductsReq{
		CategoryName: req.GetCategory(),
	})
	if err != nil {
		return nil, err
	}
	// assemble response
	var products []map[string]any
	for _, p := range resp.GetProducts() {
		products = append(products, gin.H{
			"id":          p.GetId(),
			"name":        p.GetName(),
			"description": p.GetDescription(),
			"picture":     p.GetPicture(),
			"price":       p.GetPrice(),
		})
	}
	return gin.H{
		"products": products,
	}, nil
}

func SearchProducts(c *gin.Context, req *product.SearchProductsReq) (map[string]any, error) {
	// call rpc
	resp, err := rpc.ProductClient.SearchProducts(c, &rpcproduct.SearchProductsReq{
		Query: req.GetQuery(),
	})
	if err != nil {
		return nil, err
	}
	// assemble response
	var products []map[string]any
	for _, p := range resp.GetProducts() {
		products = append(products, gin.H{
			"id":          p.GetId(),
			"name":        p.GetName(),
			"description": p.GetDescription(),
			"picture":     p.GetPicture(),
			"price":       p.GetPrice(),
		})
	}
	return gin.H{
		"query":    req.GetQuery(),
		"products": products,
	}, nil
}
