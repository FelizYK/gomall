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
	userId := getUserIdFromSession(c)
	if userId == 0 {
		return nil, nil
	}
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
	for _, item := range resp.Items {
		p, err := rpc.ProductClient.GetProduct(c, &rpcproduct.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil || p.Product == nil {
			continue
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
	userId := getUserIdFromSession(c)
	if userId == 0 {
		return errors.New("user not login")
	}
	// call rpc
	_, err = rpc.CartClient.AddCart(c, &rpccart.AddCartReq{
		UserId: userId,
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.Quantity,
		},
	})
	return
}
