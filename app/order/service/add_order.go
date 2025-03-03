package service

import (
	"context"
	"errors"

	"github.com/FelizYK/gomall/app/order/mq"
	"github.com/FelizYK/gomall/app/order/repository"
	"github.com/FelizYK/gomall/app/order/rpc"
	rpccart "github.com/FelizYK/gomall/rpc/cart"
	rpcemail "github.com/FelizYK/gomall/rpc/email"
	rpcorder "github.com/FelizYK/gomall/rpc/order"
	"github.com/FelizYK/gomall/rpc/product"
	"google.golang.org/protobuf/proto"
)

func AddOrder(ctx context.Context, req *rpcorder.AddOrderReq) (err error) {
	// get cart_items by user_id
	cartResp, err := rpc.CartClient.GetCart(ctx, &rpccart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return
	}
	if cartResp == nil || len(cartResp.Items) == 0 {
		err = errors.New("cart is empty")
		return
	}
	// assemble order
	var orderItems []repository.OrderItem
	var totalCost float32
	for _, item := range cartResp.Items {
		// get product by product_id
		productResp, errr := rpc.ProductClient.GetProduct(ctx, &product.GetProductReq{
			Id: item.ProductId,
		})
		if errr != nil {
			err = errr
			return
		}
		if productResp.Product == nil {
			err = errors.New("product not found")
			return
		}
		cost := float32(item.Quantity) * productResp.Product.Price
		orderItems = append(orderItems, repository.OrderItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Cost:      cost,
		})
		totalCost += cost
	}
	order := &repository.Order{
		UserId: req.UserId,
		Consignee: repository.Consignee{
			Email:     req.Consignee.Email,
			Firstname: req.Consignee.Firstname,
			Lastname:  req.Consignee.Lastname,
			Street:    req.Consignee.Street,
			City:      req.Consignee.City,
			Province:  req.Consignee.Province,
			Country:   req.Consignee.Country,
		},
		CreditCard: repository.CreditCard{
			CardNum:         req.CreditCard.CardNum,
			ExpirationMonth: req.CreditCard.ExpirationMonth,
			ExpirationYear:  req.CreditCard.ExpirationYear,
			Cvv:             req.CreditCard.Cvv,
		},
		TotalCost:  totalCost,
		OrderItems: orderItems,
	}
	q := repository.NewOrderQuery(ctx)
	err = q.AddOrder(order)
	if err != nil {
		return
	}
	// empty cart
	_, err = rpc.CartClient.EmptyCart(ctx, &rpccart.EmptyCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return
	}
	// send email
	data, err := proto.Marshal(&rpcemail.SendEmailReq{
		From:    "from@example.cim",
		To:      req.Consignee.Email,
		Subject: "Order Confirmation",
		Body:    "Your order has been confirmed",
	})
	if err != nil {
		return
	}
	err = mq.Nc.Publish("email", data)
	if err != nil {
		return
	}

	return nil
}
