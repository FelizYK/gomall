package handler

import (
	"net/http"

	"github.com/FelizYK/gomall/app/frontend/rpc/auth"
	"github.com/FelizYK/gomall/app/frontend/service"
	"github.com/FelizYK/gomall/app/frontend/utils"
	"github.com/gin-gonic/gin"
)

// POST /auth/register
func Register(c *gin.Context) {
	// bind form to request
	var req auth.RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// call service
	if err := service.Register(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// redirect
	c.Redirect(http.StatusFound, "/")
}

// POST /auth/login
func Login(c *gin.Context) {
	// bind form to request
	var req auth.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// call service
	if err := service.Login(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// redirect
	redirect := "/"
	if utils.ValidateNext(req.Next) {
		redirect = req.Next
	}
	c.Redirect(http.StatusFound, redirect)
}

// POST /auth/logout
func Logout(c *gin.Context) {
	// call service
	if err := service.Logout(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// redirect
	c.Redirect(http.StatusFound, "/")
}
