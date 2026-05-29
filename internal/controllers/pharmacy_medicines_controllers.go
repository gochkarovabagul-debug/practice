package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func PharmacyMedicineList(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	list, err := repositories.PharmacyMedicineList(c, repositories.PharmacyMedicineFilter{
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
func CreatePharmacyMedicine(c *gin.Context) {
	var req models.PharmacyMedicinesCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = repositories.CreatePharmacyMedicine(c.Request.Context(), req.Name, req.Description, req.Price, req.NewPrice, req.CategoryId)
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(200, gin.H{
	// 	"success": true,
	// })
	utils.SuccessResponse(c, "medicine created")
}
func DeletePharmacyMedicine(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err := repositories.DeletePharmacyMedicine(c.Request.Context(), id)
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"success": true,
	// })
	utils.SuccessResponse(c, "medicine deleted")
}
func UpdatePharmacyMedicine(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	var req models.PharmacyMedicinesCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = repositories.UpdatePharmacyMedicine(c.Request.Context(), id, req)
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"success": true,
	// })
	utils.SuccessResponse(c, "medicine updated")
}
func GetPharmacyMedicine(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	req, err := repositories.GetPharmacyMedicine(c.Request.Context(), id)
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"success": true,
	// 	"data":    req,
	// })
	utils.SuccessResponse(c, req)
}

func PharmacyMedicinesRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/pharmacymedicines", PharmacyMedicineList)
	rg.POST("/admin/pharmacymedicines/create", CreatePharmacyMedicine)
	rg.DELETE("/admin/pharmacymedicines/delete/:id", DeletePharmacyMedicine)
	rg.GET("/admin/pharmacymedicines/get/:id", GetPharmacyMedicine)
	rg.PUT("/admin/pharmacymedicines/update/:id", UpdatePharmacyMedicine)
}
