package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul/practice/internal/models"
	"github.com/gochkarovabagul/practice/internal/repositories"
)

func CreateCategory(c *gin.Context) {
	var req models.CategoryCreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err1 := repositories.CreateCategory(c.Request.Context(), req.Name)
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

func CategoryRoutes(rg *gin.RouterGroup) {
	rg.POST("/admin/categories", CreateCategory)
}
