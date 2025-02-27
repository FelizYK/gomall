package router

import (
	"github.com/FelizYK/gomall/app/frontend/handler"
	"github.com/gin-gonic/gin"
)

func Checkout(r *gin.Engine) {
	// checkout
	r.GET("/cart", handler.CheckoutPage)
	// pay
	r.POST("/cart", handler.Checkout)
}
