package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/services"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func OrderList(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	list, err := services.OrderListService(c, models.OrderFilter{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, list)
}
func CreateOrder(c *gin.Context) {
	var req models.OrderCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = services.CreateOrderService(c, req.Name, req.Price, req.Description)

	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "order created")
}
func DeleteOrder(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err := services.DeleteOrderService(c, id)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "order deleted")
}
func UpdateOrder(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	var req models.OrderCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = services.UpdateOrderService(c, id, req)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "order updated")
}
func GetOrder(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	req, err := services.GetOrderServices(c, id)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, req)
}

func OrderRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/orders", OrderList)
	rg.POST("/admin/orders/create", CreateOrder)
	rg.DELETE("/admin/orders/delete/:id", DeleteOrder)
	rg.GET("/admin/orders/get/:id", GetOrder)
	rg.PUT("/admin/orders/update/:id", UpdateOrder)
}
