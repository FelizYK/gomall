package router

import (
	"github.com/FelizYK/gomall/app/frontend/handler"
	"github.com/FelizYK/gomall/app/frontend/middleware"
	"github.com/gin-gonic/gin"
)

func Checkout(r *gin.Engine) {
	protected := r.Group("/checkout", middleware.Auth())
	{
		// checkout
		protected.GET("/", handler.CheckoutPage)
		// pay
		protected.POST("/", handler.Checkout)
	}
}
