package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
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
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(200, gin.H{
	// 	"success": true,
	// 	"data":    list,
	// })
	utils.SuccessResponse(c, list)
}
func CreatePharmacy(c *gin.Context) {
	var req models.PharmacyCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = repositories.CreatePharmacy(c.Request.Context(), req.Name, req.Address, req.Hours)
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(200, gin.H{
	// 	"success": true,
	// })
	utils.SuccessResponse(c, "pharmacy created")
}
func DeletePharmacy(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err := repositories.DeletePharmacy(c.Request.Context(), id)
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(200, gin.H{
	// 	"success": true,
	// })
	utils.SuccessResponse(c, "pharmacy deleted")
}
func GetPharmacy(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	req, err := repositories.GetPharmacy(c.Request.Context(), id)
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(200, gin.H{
	// 	"success": true,
	// 	"data":    req,
	// })
	utils.SuccessResponse(c, req)
}
func UpdatePharmacy(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	var req models.PharmacyCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = repositories.UpdatePharmacy(c.Request.Context(), id, req)
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(200, gin.H{
	// 	"success": true,
	// })
	utils.SuccessResponse(c, "pharmacy updated")
}
func FindNearbyPharmacies(c *gin.Context) {
	latstr := c.Query("latitude")
	lat, _ := strconv.ParseFloat(latstr, 64)
	lonstr := c.Query("longitude")
	lon, _ := strconv.ParseFloat(lonstr, 64)
	result, err := repositories.FindNearbyPharmacies(c.Request.Context(), lon, lat)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, result)
}
func PharmacyRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/phatmacies/findnearpharmacy", FindNearbyPharmacies)
	rg.GET("/admin/pharmacies", PharmacyList)
	rg.POST("/admin/pharmacies/create", CreatePharmacy)
	rg.DELETE("/admin/pharmacies/delete/:id", DeletePharmacy)
	rg.GET("/admin/pharmacies/get/:id", GetPharmacy)
	rg.PUT("/admin/pharmacies/update/:id", UpdatePharmacy)
}
