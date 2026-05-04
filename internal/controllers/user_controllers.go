package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul/internal/repositories"
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
	return
}
func UserRoutes(rg *gin.RouterGroup) {
	rg.GET("admin/users", UserList)
}
