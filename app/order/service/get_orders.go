package service

import (
	"context"

	"github.com/FelizYK/gomall/app/order/repository"
	rpcorder "github.com/FelizYK/gomall/rpc/order"
)

func GetOrders(ctx context.Context, req *rpcorder.GetOrdersReq) (resp *rpcorder.GetOrdersResp, err error) {
	// get orders by user_id
	q := repository.NewOrderQuery(ctx)
	orders, err := q.GetOrders(req.UserId)
	if err != nil {
		return
	}
	// assemble response
	resp = &rpcorder.GetOrdersResp{}
	for _, order := range orders {
		var items []*rpcorder.OrderItem
		for _, item := range order.OrderItems {
			items = append(items, &rpcorder.OrderItem{
				ProductId: item.ProductId,
				Quantity:  item.Quantity,
				Cost:      item.Cost,
			})
		}
		resp.Orders = append(resp.Orders, &rpcorder.Order{
			Id:        uint32(order.ID),
			CreatedAt: order.CreatedAt.Format("2006-01-02 15:04:05"),
			Consignee: &rpcorder.Consignee{
				Email:     order.Consignee.Email,
				Firstname: order.Consignee.Firstname,
				Lastname:  order.Consignee.Lastname,
				Street:    order.Consignee.Street,
				City:      order.Consignee.City,
				Province:  order.Consignee.Province,
				Country:   order.Consignee.Country,
			},
			TotalCost:  order.TotalCost,
			OrderItems: items,
		})
	}
	return
}
