package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/permission"
	"github.com/gochkarovabagul-debug/practice/internal/services"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func PharmacyList(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	list, err := services.PharmacyListService(c, models.PharmacyFilter{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, list)
}
func CreatePharmacy(c *gin.Context) {
	var req models.PharmacyCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = services.CreatePharmacyService(c, req.Name, req.Address, req.Hours, req.AdminUserId)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "pharmacy created")
}
func DeletePharmacy(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err := services.DeletePharmacyService(c, id)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "pharmacy deleted")
}
func GetPharmacy(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	req, err := services.GetPharmacyService(c, id)
	if utils.ErrorCheck(c, err) {
		return
	}
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
	err = services.UpdatePharmacyService(c, id, req)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "pharmacy updated")
}
func FindNearbyPharmacies(c *gin.Context) {
	latstr := c.Query("latitude")
	lat, _ := strconv.ParseFloat(latstr, 64)
	lonstr := c.Query("longitude")
	lon, _ := strconv.ParseFloat(lonstr, 64)
	result, err := services.FindNearbyPharmaciesService(c.Request.Context(), lon, lat)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, result)
}
func PharmacyRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/pharmacies/findnearpharmacy", FindNearbyPharmacies)
	rg.GET("/admin/pharmacies", PharmacyList)
	rg.POST("/admin/pharmacies/create", CreatePharmacy).Use(permission.RequirePharmacyAdmin())
	rg.DELETE("/admin/pharmacies/delete/:id", DeletePharmacy).Use(permission.RequirePharmacyAdmin())
	rg.GET("/admin/pharmacies/get/:id", GetPharmacy)
	rg.PUT("/admin/pharmacies/update/:id", UpdatePharmacy).Use(permission.RequirePharmacyAdmin())
}
