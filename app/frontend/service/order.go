package service

import (
	"errors"

	"github.com/FelizYK/gomall/app/frontend/rpc"
	rpcorder "github.com/FelizYK/gomall/rpc/order"
	rpcproduct "github.com/FelizYK/gomall/rpc/product"
	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) (map[string]any, error) {
	// get user_id
	userId, exists := c.Get("user_id")
	if !exists {
		return nil, errors.New("user not login")
	}
	// call rpc
	resp, err := rpc.OrderClient.GetOrders(c, &rpcorder.GetOrdersReq{
		UserId: userId.(uint32),
	})
	if err != nil {
		return nil, err
	}
	// assemble response
	if resp == nil || len(resp.Orders) == 0 {
		return gin.H{}, nil
	}
	var orders []map[string]any
	for _, order := range resp.Orders {
		var products []map[string]any
		for _, item := range order.OrderItems {
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
				"cost":     item.Cost,
			})
		}
		orders = append(orders, gin.H{
			"id":         order.Id,
			"created_at": order.CreatedAt,
			"total_cost": order.TotalCost,
			"products":   products,
		})
	}
	return gin.H{
		"orders": orders,
	}, nil
}
