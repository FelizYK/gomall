package service

import (
	"errors"

	"github.com/FelizYK/gomall/app/frontend/rpc"
	"github.com/FelizYK/gomall/app/frontend/rpc/checkout"
	rpccheckout "github.com/FelizYK/gomall/rpc/checkout"
	"github.com/gin-gonic/gin"
)

func Checkout(c *gin.Context, req *checkout.CheckoutReq) (err error) {
	// get user_id
	userId := getUserIdFromSession(c)
	if userId == 0 {
		return errors.New("user not login")
	}
	// call rpc
	_, err = rpc.CheckoutClient.Checkout(c, &rpccheckout.CheckoutReq{
		UserId: userId,
		Consignee: &rpccheckout.Consignee{
			Email:     req.Email,
			Firstname: req.Firstname,
			Lastname:  req.Lastname,
			Street:    req.Street,
			City:      req.City,
			Province:  req.Province,
			Country:   req.Country,
		},
		CreditCard: &rpccheckout.CreditCard{
			CardNum:         req.CardNum,
			ExpirationYear:  req.ExpirationYear,
			ExpirationMonth: req.ExpirationMonth,
			Cvv:             req.Cvv,
		},
	})
	return
}
