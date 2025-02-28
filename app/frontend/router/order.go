package router

import (
	"github.com/FelizYK/gomall/app/frontend/handler"
	"github.com/gin-gonic/gin"
)

func Order(r *gin.Engine) {
	// order
	r.GET("/order", handler.OrderPage)
}
