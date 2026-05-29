package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
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
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(200, gin.H{
	// 	"success": true,
	// 	"data":    list,
	// })
	utils.SuccessResponse(c, list)
}
func CreateCategory(c *gin.Context) {
	var req models.CategoryCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = repositories.CreateCategory(c.Request.Context(), req.Name)
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(200, gin.H{
	// 	"success": true,
	// })
	utils.SuccessResponse(c, "category created")
}
func DeleteCategory(c *gin.Context) {
	categoryidstr := c.Param("id")
	categoryid, _ := strconv.Atoi(categoryidstr)
	err := repositories.DeleteCategory(c.Request.Context(), categoryid)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "category deleted")
	// c.JSON(http.StatusOK, gin.H{
	// 	"success": true,
	// })
}
func UpdateCategory(c *gin.Context) {
	categoryidstr := c.Param("id")
	categoryid, _ := strconv.Atoi(categoryidstr)
	var req models.CategoryCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = repositories.UpdateCategory(c.Request.Context(), categoryid, req)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "category updated")
	// c.JSON(http.StatusOK, gin.H{
	// 	"success": true,
	// })
}
func GetCategory(c *gin.Context) {
	categoryidstr := c.Param("id")
	categoryid, _ := strconv.Atoi(categoryidstr)
	req, err := repositories.GetCategory(c.Request.Context(), categoryid)
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"success": true,
	// 	"data":    req,
	// })
	utils.SuccessResponse(c, req)
}

func CategoryRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/categories", CategoryList)
	rg.POST("/admin/categories/create", CreateCategory)
	rg.DELETE("/admin/categories/delete/:categoryid", DeleteCategory)
	rg.GET("/admin/categories/get/:categoryid", GetCategory)
	rg.PUT("/admin/categories/update/:categoryid", UpdateCategory)
}
