package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func WrapResponse(c *gin.Context, resp map[string]any) map[string]any {
	resp["user_id"] = GetUserIdFromSession(c)
	return resp
}

func GetUserIdFromSession(c *gin.Context) interface{} {
	session := sessions.Default(c)
	return session.Get("user_id")
}
