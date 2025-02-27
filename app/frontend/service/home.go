package service

import (
	"github.com/FelizYK/gomall/app/frontend/rpc"
	rpcproduct "github.com/FelizYK/gomall/rpc/product"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) (map[string]any, error) {
	// call rpc
	resp, err := rpc.ProductClient.ListProducts(c, &rpcproduct.ListProductsReq{})
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
		"products_num": len(products),
		"products":     products,
	}, nil
}
