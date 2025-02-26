package service

import (
	"github.com/FelizYK/gomall/app/frontend/rpc"
	"github.com/FelizYK/gomall/app/frontend/rpc/cart"
	rpccart "github.com/FelizYK/gomall/rpc/cart"
	rpcproduct "github.com/FelizYK/gomall/rpc/product"
	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context, userId uint32) (map[string]any, error) {
	// call rpc
	resp, err := rpc.CartClient.GetCart(c, &rpccart.GetCartReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	// assemble response
	var products []map[string]any
	var total float32
	for _, p := range resp.GetProducts() {
		presp, err := rpc.ProductClient.GetProduct(c, &rpcproduct.GetProductReq{
			Id: p.GetProductId(),
		})
		if err != nil || presp.GetProduct() == nil {
			continue
		}
		product := presp.GetProduct()
		products = append(products, gin.H{
			"name":     product.GetName(),
			"picture":  product.GetPicture(),
			"price":    product.GetPrice(),
			"quantity": p.GetQuantity(),
		})
		total += product.GetPrice() * float32(p.GetQuantity())
	}
	return gin.H{
		"products": products,
		"total":    total,
	}, nil
}

func AddCart(c *gin.Context, req *cart.AddCartReq, userId uint32) (err error) {
	// call rpc
	_, err = rpc.CartClient.AddCart(c, &rpccart.AddCartReq{
		UserId: userId,
		Item: rpccart.CartItem{
			ProductId: req.GetProductId(),
			Quantity:  req.GetQuantity(),
		},
	})
	return
}
