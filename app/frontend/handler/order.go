package handler

import (
	"net/http"

	"github.com/FelizYK/gomall/app/frontend/service"
	"github.com/gin-gonic/gin"
)

// GET /order
func OrderPage(c *gin.Context) {
	// call service to get orders
	resp, err := service.GetOrders(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// render html
	c.HTML(http.StatusOK, "order.html", service.WrapResponse(c, resp))
}
