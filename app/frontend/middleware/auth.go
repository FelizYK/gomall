package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GlobalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := getUserIdFromSession(c)
		if userId != 0 {
			c.Set("user_id", userId)
		}
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := getUserIdFromSession(c)
		if userId == 0 {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Set("user_id", userId)
		c.Next()
	}
}

func getUserIdFromSession(c *gin.Context) uint32 {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	if userId == nil {
		return 0
	}
	return userId.(uint32)
}
