package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func WrapResponse(c *gin.Context, resp map[string]any) map[string]any {
	session := sessions.Default(c)
	resp["user_id"] = session.Get("user_id")
	return resp
}
