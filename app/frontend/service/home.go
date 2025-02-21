package service

import (
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) (map[string]any, error) {
	items := []map[string]any{
		{"Name": "t-shirt-1", "Price": 100, "Picture": "/assets/t-shirt-1.jpeg"},
		{"Name": "t-shirt-2", "Price": 200, "Picture": "/assets/t-shirt-2.jpeg"},
		{"Name": "t-shirt-3", "Price": 300, "Picture": "/assets/t-shirt-3.jpeg"},
		{"Name": "t-shirt-1", "Price": 100, "Picture": "/assets/t-shirt-1.jpeg"},
		{"Name": "t-shirt-2", "Price": 200, "Picture": "/assets/t-shirt-2.jpeg"},
		{"Name": "t-shirt-3", "Price": 300, "Picture": "/assets/t-shirt-3.jpeg"},
	}
	return gin.H{
		"cart_num": 1,
		"items":    items,
	}, nil
}
