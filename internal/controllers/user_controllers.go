package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
)

func UserList(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	role := c.Query("role")
	list, err := repositories.UserList(c, repositories.UserFilter{
		Limit:  limit,
		Offset: offset,
		Search: search,
		Role:   role,
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
func CreateUser(c *gin.Context) {
	var req models.UserCreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err1 := repositories.CreateUser(c.Request.Context(), req.FirstName, req.LastName, req.Role, req.Password, req.Email)
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
func DeleteUser(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err1 := repositories.DeleteUser(c.Request.Context(), id)
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
func GetUser(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	// var req models.UserResponse
	req, err1 := repositories.GetUser(c.Request.Context(), id)
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
func UpdateUser(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	// var req models.UserResponse
	var req models.UserCreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = repositories.UpdateUser(c.Request.Context(), id, req)
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
func UserRoutes(rg *gin.RouterGroup) {
	rg.GET("/admin/users", UserList)
	rg.POST("/admin/users/create", CreateUser)
	rg.DELETE("/admin/users/delete/:id", DeleteUser)
	rg.GET("/admin/users/get/:id", GetUser)
	rg.PUT("/admin/users/update/:id", UpdateUser)
}
