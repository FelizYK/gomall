package handler

import (
	"net/http"

	"github.com/FelizYK/gomall/app/frontend/rpc/cart"
	"github.com/FelizYK/gomall/app/frontend/service"
	"github.com/gin-gonic/gin"
)

// GET /cart
func GetCart(c *gin.Context) {
	// call service
	resp, err := service.GetCart(c, GetUserIdFromSession(c).(uint32))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// render html
	c.HTML(http.StatusOK, "cart.html", WrapResponse(c, resp))
}

// POST /cart
func AddCart(c *gin.Context) {
	var req cart.AddCartReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// call service
	if err := service.AddCart(c, &req, GetUserIdFromSession(c).(uint32)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// redirect
	c.Redirect(http.StatusFound, "/cart")
}
