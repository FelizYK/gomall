package middleware

import (
	"net/http"

	"github.com/FelizYK/gomall/app/frontend/service"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := service.GetUserIdFromSession(c)
		if userId == 0 {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Set("user_id", userId)
		c.Next()
	}
}
