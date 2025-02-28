package router

import (
	"github.com/FelizYK/gomall/app/frontend/handler"
	"github.com/gin-gonic/gin"
)

func Checkout(r *gin.Engine) {
	// checkout
	r.GET("/checkout", handler.CheckoutPage)
	// pay
	r.POST("/checkout", handler.Checkout)
}
