package service

import (
	"errors"
	"strconv"

	"github.com/FelizYK/gomall/app/frontend/rpc"
	"github.com/FelizYK/gomall/app/frontend/rpc/cart"
	rpccart "github.com/FelizYK/gomall/rpc/cart"
	rpcproduct "github.com/FelizYK/gomall/rpc/product"
	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) (map[string]any, error) {
	// get user_id
	userId, exists := c.Get("user_id")
	if !exists {
		return nil, errors.New("user not login")
	}
	// call rpc
	resp, err := rpc.CartClient.GetCart(c, &rpccart.GetCartReq{
		UserId: userId.(uint32),
	})
	if err != nil {
		return nil, err
	}
	// assemble response
	var products []map[string]any
	var total float32
	for _, item := range resp.Items {
		p, err := rpc.ProductClient.GetProduct(c, &rpcproduct.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			return nil, err
		}
		products = append(products, gin.H{
			"name":     p.Product.Name,
			"picture":  p.Product.Picture,
			"price":    p.Product.Price,
			"quantity": item.Quantity,
		})
		total += p.Product.Price * float32(item.Quantity)
	}
	return gin.H{
		"products": products,
		"total":    strconv.FormatFloat(float64(total), 'f', 2, 32),
	}, nil
}

func AddCart(c *gin.Context, req *cart.AddCartReq) (err error) {
	// get user_id
	userId, exists := c.Get("user_id")
	if !exists {
		return errors.New("user not login")
	}
	// call rpc
	_, err = rpc.CartClient.AddCart(c, &rpccart.AddCartReq{
		UserId: userId.(uint32),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.Quantity,
		},
	})
	return
}
