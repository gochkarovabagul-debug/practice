package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
)

func OrderList(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	list, err := repositories.OrderList(c, repositories.OrderFilter{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"success":   false,
			"error_msg": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"data":    list,
	})
}
func CreateOrder(c *gin.Context) {
	var req models.OrderCreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err1 := repositories.CreateOrder(c.Request.Context(), req.Name, req.Price, req.Description)
	if err1 != nil {
		c.JSON(500, gin.H{
			"error": err1.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
	})
}
func DeleteOrder(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err1 := repositories.DeleteOrder(c.Request.Context(), id)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err1.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
func UpdateOrder(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	var req models.OrderCreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = repositories.UpdateOrder(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
func GetOrder(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	req, err1 := repositories.GetOrder(c.Request.Context(), id)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err1.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    req,
	})
}

func OrderRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/orders", OrderList)
	rg.POST("/admin/orders/create", CreateOrder)
	rg.DELETE("/admin/orders/delete/:id", DeleteOrder)
	rg.GET("/admin/orders/get/:id", GetOrder)
	rg.PUT("/admin/orders/update/:id", UpdateOrder)
}
