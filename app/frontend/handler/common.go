package handler

import (
	"log"

	"github.com/FelizYK/gomall/app/frontend/rpc"
	rpccart "github.com/FelizYK/gomall/rpc/cart"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func WrapResponse(c *gin.Context, resp map[string]any) map[string]any {
	// user_id
	userId := GetUserIdFromSession(c)
	if userId != 0 {
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

func GetUserIdFromSession(c *gin.Context) uint32 {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	if userId == nil {
		return 0
	}
	return userId.(uint32)
}
