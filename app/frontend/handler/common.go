package handler

import (
	"log"

	"github.com/FelizYK/gomall/app/frontend/rpc"
	rpccart "github.com/FelizYK/gomall/rpc/cart"
	"github.com/gin-gonic/gin"
)

func WrapResponse(c *gin.Context, resp map[string]any) map[string]any {
	// user_id
	userId, exists := c.Get("user_id")
	if exists {
		userId := userId.(uint32)
		resp["user_id"] = userId

		// cart_num
		cartResp, err := rpc.CartClient.GetCart(c, &rpccart.GetCartReq{
			UserId: userId,
		})
		if err != nil {
			log.Fatalf("call rpc.CartClient.GetCart failed, err: %v", err)
		}
		cartNum := len(cartResp.Items)
		resp["cart_num"] = cartNum
	}

	return resp
}
