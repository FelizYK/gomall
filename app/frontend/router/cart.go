package router

import (
	"github.com/FelizYK/gomall/app/frontend/handler"
	"github.com/FelizYK/gomall/app/frontend/middleware"
	"github.com/gin-gonic/gin"
)

func Cart(r *gin.Engine) {
	protected := r.Group("/cart", middleware.Auth())
	{
		// cart
		protected.GET("/cart", handler.GetCart)
		// add product to cart
		protected.POST("/cart", handler.AddCart)
	}
}
