package router

import (
	"github.com/FelizYK/gomall/app/frontend/handler"
	"github.com/gin-gonic/gin"
)

func Cart(r *gin.Engine) {
	// cart
	r.GET("/cart", handler.GetCart)
	// add product to cart
	r.POST("/cart", handler.AddCart)
}
