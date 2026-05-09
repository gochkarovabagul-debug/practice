package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
)

func CategoryList(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	list, err := repositories.CategoryList(c, repositories.CategoryFilter{
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
func DeleteCategory(c *gin.Context) {
	categoryidstr := c.Param("id")
	categoryid, _ := strconv.Atoi(categoryidstr)
	err1 := repositories.DeleteCategory(c.Request.Context(), categoryid)
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
func UpdateCategory(c *gin.Context) {
	categoryidstr := c.Param("id")
	categoryid, _ := strconv.Atoi(categoryidstr)
	var req models.CategoryCreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = repositories.UpdateCategory(c.Request.Context(), categoryid, req)
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
func GetCategory(c *gin.Context) {
	categoryidstr := c.Param("id")
	categoryid, _ := strconv.Atoi(categoryidstr)
	req, err1 := repositories.GetCategory(c.Request.Context(), categoryid)
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

func CategoryRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/categories", CategoryList)
	rg.POST("/admin/categories/create", CreateCategory)
	rg.DELETE("/admin/categories/delete/:categoryid", DeleteCategory)
	rg.GET("/admin/categories/get/:categoryid", GetCategory)
	rg.PUT("/admin/categories/update/:categoryid", UpdateCategory)
}
