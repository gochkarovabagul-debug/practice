package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
)

func PharmacyList(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	list, err := repositories.PharmacyList(c, repositories.PharmacyFilter{
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
func CreatePharmacy(c *gin.Context) {
	var req models.PharmacyCreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err1 := repositories.CreatePharmacy(c.Request.Context(), req.Name, req.Address, req.Hours)
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
func DeletePharmacy(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err1 := repositories.DeletePharmacy(c.Request.Context(), id)
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
func GetPharmacy(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	// var req models.UserResponse
	req, err1 := repositories.GetPharmacy(c.Request.Context(), id)
	if err1 != nil {
		c.JSON(500, gin.H{
			"error": err1.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"data":    req,
	})
}
func UpdatePharmacy(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	// var req models.UserResponse
	var req models.PharmacyCreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = repositories.UpdatePharmacy(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
	})
}
func PharmacyRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/pharmacies", PharmacyList)
	rg.POST("/admin/pharmacies/create", CreatePharmacy)
	rg.DELETE("/admin/pharmacies/delete/:id", DeletePharmacy)
	rg.GET("/admin/pharmacies/get/:id", GetPharmacy)
	rg.PUT("/admin/pharmacies/update/:id", UpdatePharmacy)
}
