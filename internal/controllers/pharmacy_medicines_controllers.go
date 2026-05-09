package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
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
func CreatePharmacyMedicine(c *gin.Context) {
	var req models.PharmacyMedicinesCreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err1 := repositories.CreatePharmacyMedicine(c.Request.Context(), req.Name, req.Description, req.Price, req.NewPrice, req.CategoryId)
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
func DeletePharmacyMedicine(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err1 := repositories.DeletePharmacyMedicine(c.Request.Context(), id)
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
func UpdatePharmacyMedicine(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	var req models.PharmacyMedicinesCreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = repositories.UpdatePharmacyMedicine(c.Request.Context(), id, req)
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
func GetPharmacyMedicine(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	req, err1 := repositories.GetPharmacyMedicine(c.Request.Context(), id)
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

func PharmacyMedicinesRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/pharmacymedicines", PharmacyMedicineList)
	rg.POST("/admin/pharmacymedicines/create", CreatePharmacyMedicine)
	rg.DELETE("/admin/pharmacymedicines/delete/:id", DeletePharmacyMedicine)
	rg.GET("/admin/pharmacymedicines/get/:id", GetPharmacyMedicine)
	rg.PUT("/admin/pharmacymedicines/update/:id", UpdatePharmacyMedicine)
}
