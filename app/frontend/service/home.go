package service

import (
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) (map[string]any, error) {
	items := []map[string]any{
		{"name": "t-shirt-1", "price": 100, "picture": "/assets/t-shirt-1.jpeg"},
		{"name": "t-shirt-2", "price": 200, "picture": "/assets/t-shirt-2.jpeg"},
		{"name": "t-shirt-3", "price": 300, "picture": "/assets/t-shirt-3.jpeg"},
		{"name": "t-shirt-1", "price": 100, "picture": "/assets/t-shirt-1.jpeg"},
		{"name": "t-shirt-2", "price": 200, "picture": "/assets/t-shirt-2.jpeg"},
		{"name": "t-shirt-3", "price": 300, "picture": "/assets/t-shirt-3.jpeg"},
	}
	return gin.H{
		"cart_num": 1,
		"items":    items,
	}, nil
}
