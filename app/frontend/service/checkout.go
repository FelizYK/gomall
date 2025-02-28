package service

import (
	"errors"

	"github.com/FelizYK/gomall/app/frontend/rpc"
	"github.com/FelizYK/gomall/app/frontend/rpc/checkout"
	rpcorder "github.com/FelizYK/gomall/rpc/order"
	"github.com/gin-gonic/gin"
)

func Checkout(c *gin.Context, req *checkout.CheckoutReq) (err error) {
	// get user_id
	userId := getUserIdFromSession(c)
	if userId == 0 {
		return errors.New("user not login")
	}
	// call rpc
	_, err = rpc.OrderClient.AddOrder(c, &rpcorder.AddOrderReq{
		UserId: userId,
		Consignee: &rpcorder.Consignee{
			Email:     req.Email,
			Firstname: req.Firstname,
			Lastname:  req.Lastname,
			Street:    req.Street,
			City:      req.City,
			Province:  req.Province,
			Country:   req.Country,
		},
		CreditCard: &rpcorder.CreditCard{
			CardNum:         req.CardNum,
			ExpirationYear:  req.ExpirationYear,
			ExpirationMonth: req.ExpirationMonth,
			Cvv:             req.Cvv,
		},
	})
	return
}
