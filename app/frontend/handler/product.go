package handler

import (
	"net/http"

	"github.com/FelizYK/gomall/app/frontend/rpc/product"
	"github.com/FelizYK/gomall/app/frontend/service"
	"github.com/gin-gonic/gin"
)

// GET /product
func GetProduct(c *gin.Context) {
	// bind query to request
	var req product.GetProductReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// call service
	resp, err := service.GetProduct(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// render html
	c.HTML(http.StatusOK, "product.html", WrapResponse(c, resp))
}

// GET /category/:category
func ListProductsByCategory(c *gin.Context) {
	// bind path to request
	var req product.ListProductsReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// call service
	resp, err := service.ListProductsByCategory(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// render html
	c.HTML(http.StatusOK, "category.html", WrapResponse(c, resp))
}

// GET /search
func SearchProducts(c *gin.Context) {
	// bind query to request
	var req product.SearchProductsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// call service
	resp, err := service.SearchProducts(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// render html
	c.HTML(http.StatusOK, "search.html", WrapResponse(c, resp))
}
