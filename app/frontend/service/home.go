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
		"products_num": len(products),
		"products":     products,
	}, nil
}
