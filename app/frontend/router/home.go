package router

import (
	"github.com/FelizYK/gomall/app/frontend/handler"
	"github.com/gin-gonic/gin"
)

func Home(r *gin.Engine) {
	r.GET("/", handler.Home)
}
