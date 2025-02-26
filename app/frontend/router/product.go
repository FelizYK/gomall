package router

import (
	"github.com/FelizYK/gomall/app/frontend/handler"
	"github.com/gin-gonic/gin"
)

func Product(r *gin.Engine) {
	// category
	r.GET("/category/:category", handler.ListProductsByCategory)
	// product
	r.GET("/product", handler.GetProduct)
	// search
	r.GET("/search", handler.SearchProducts)
}
