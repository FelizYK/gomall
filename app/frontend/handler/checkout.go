package handler

import (
	"net/http"

	"github.com/FelizYK/gomall/app/frontend/rpc/checkout"
	"github.com/FelizYK/gomall/app/frontend/service"
	"github.com/gin-gonic/gin"
)

// GET /checkout
func CheckoutPage(c *gin.Context) {
	// call cart service to get items
	resp, err := service.GetCart(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// render html
	c.HTML(http.StatusOK, "checkout.html", WrapResponse(c, resp))
}

// POST /checkout
func Checkout(c *gin.Context) {
	// bind form to request
	var req checkout.CheckoutReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// call service
	if err := service.Checkout(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// redirect
	c.Redirect(http.StatusFound, "/order")
}
