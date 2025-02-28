package router

import (
	"net/http"

	"github.com/FelizYK/gomall/app/frontend/handler"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {
	// sign up
	r.GET("/register", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "register.html", nil)
	})
	r.POST("/auth/register", handler.Register)

	// sign in
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/auth/login", handler.Login)

	// sign out
	r.POST("/auth/logout", handler.Logout)
}
