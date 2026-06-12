package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/permission"
	"github.com/gochkarovabagul-debug/practice/internal/services"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func CategoryList(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	list, err := services.CategoryListService(c, models.CategoryFilter{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, list)
}
func CreateCategory(c *gin.Context) {
	var req models.CategoryCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = services.CreateCategoryService(c, req.Name)
	if utils.ErrorCheck(c, err) {
		return
	}

	utils.SuccessResponse(c, "category created")
}
func DeleteCategory(c *gin.Context) {
	categoryidstr := c.Param("id")
	categoryid, _ := strconv.Atoi(categoryidstr)
	err := services.DeleteCategoryService(c, categoryid)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "category deleted")
}
func UpdateCategory(c *gin.Context) {
	categoryidstr := c.Param("id")
	categoryid, _ := strconv.Atoi(categoryidstr)
	var req models.CategoryCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = services.UpdateCategoryService(c, categoryid, req)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "category updated")

}
func GetCategory(c *gin.Context) {
	categoryidstr := c.Param("id")
	categoryid, _ := strconv.Atoi(categoryidstr)
	req, err := services.GetCategoryService(c, categoryid)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, req)
}

func CategoryRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/categories", CategoryList)
	rg.POST("/admin/categories/create", CreateCategory).Use(permission.RequireAdmin())
	rg.DELETE("/admin/categories/delete/:categoryid", DeleteCategory).Use(permission.RequireAdmin())
	rg.GET("/admin/categories/get/:categoryid", GetCategory)
	rg.PUT("/admin/categories/update/:categoryid", UpdateCategory).Use(permission.RequireAdmin())
}
