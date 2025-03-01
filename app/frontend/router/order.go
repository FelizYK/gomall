package router

import (
	"github.com/FelizYK/gomall/app/frontend/handler"
	"github.com/FelizYK/gomall/app/frontend/middleware"
	"github.com/gin-gonic/gin"
)

func Order(r *gin.Engine) {
	protected := r.Group("/order", middleware.Auth())
	{
		// order
		protected.GET("/order", handler.OrderPage)
	}
}
