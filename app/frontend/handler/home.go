package handler

import (
	"net/http"

	"github.com/FelizYK/gomall/app/frontend/service"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	resp, err := service.Home(c)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "home.html", WrapResponse(c, resp))
}
