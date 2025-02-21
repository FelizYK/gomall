package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "github.com/FelizYK/gomall/frontend/handler"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")
	r.Static("assets", "./assets")

	r.GET("/", handler.Home)

	r.Run() // listen and serve on 0.0.0.0:8080
}
