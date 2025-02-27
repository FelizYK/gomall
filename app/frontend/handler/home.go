package handler

import (
	"net/http"

	"github.com/FelizYK/gomall/app/frontend/service"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	// call service
	resp, err := service.Home(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// render html
	c.HTML(http.StatusOK, "home.html", service.WrapResponse(c, resp))
}
