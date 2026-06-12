package controllers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/permission"
	"github.com/gochkarovabagul-debug/practice/internal/services"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func PharmacyMedicineList(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	list, err := services.PharmacyMedicineListService(c, models.PharmacyMedicineFilter{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, list)
}
func CreatePharmacyMedicine(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.TrimPrefix(auth, "Bearer ")
	token = strings.TrimSpace(token)
	var req models.PharmacyMedicinesCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = services.CreatePharmacyMedicineService(c, req.Name, req.Description, req.Price, req.NewPrice, req.CategoryId, req.PharmacyId, token)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}
	utils.SuccessResponse(c, "medicine created")
}
func DeletePharmacyMedicine(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err := services.DeletePharmacyMedicineService(c, id)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}
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
	err = services.UpdatePharmacyMedicineService(c, id, req)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "medicine updated")
}
func GetPharmacyMedicine(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	req, err := services.GetPharmacyMedicineServices(c, id)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, req)
}
func PharmacyMedicinesRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/pharmacymedicines", PharmacyMedicineList)
	rg.POST("/admin/pharmacymedicines/create", CreatePharmacyMedicine).Use(permission.RequirePharmacyAdmin())
	rg.DELETE("/admin/pharmacymedicines/delete/:id", DeletePharmacyMedicine).Use(permission.RequirePharmacyAdmin())
	rg.GET("/admin/pharmacymedicines/get/:id", GetPharmacyMedicine)
	rg.PUT("/admin/pharmacymedicines/update/:id", UpdatePharmacyMedicine).Use(permission.RequirePharmacyAdmin())
}
