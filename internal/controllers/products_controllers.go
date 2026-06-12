package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func ProductList(c *gin.Context) {
	limitstr := c.Query("limit")
	limit, _ := strconv.Atoi(limitstr)
	offsetstr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetstr)
	search := c.Query("search")
	list, err := repositories.ProductList(c, models.ProductFilter{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, list)
}
func CreateProduct(c *gin.Context) {
	var product models.ProductCreateRequest
	err := c.Bind(&product)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = repositories.CreateProduct(c.Request.Context(), product.Name, product.PharmacyId)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "product created")
}
func DeleteProduct(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err := repositories.DeleteProduct(c, id)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "product deleted")
}
func UpdateProduct(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	var product models.ProductCreateRequest
	err := c.Bind(&product)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = repositories.UpdateProduct(c, id, product)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "product updated")
}
func GetProduct(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	product, err := repositories.GetProduct(c, id)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, product)
}
func ProductRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/products", ProductList)
	rg.POST("/admin/products/create", CreateProduct)
	rg.DELETE("/admin/products/delete/:id", DeleteProduct)
	rg.GET("/admin/products/get/:id", GetProduct)
	rg.PUT("/admin/products/update/:id", UpdateProduct)
}
