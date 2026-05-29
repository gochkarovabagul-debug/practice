package controllers

import (
	"crypto/rand"
	"encoding/hex"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
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

	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(200, gin.H{
	// 	"success": true,
	// 	"data":    list,
	// })
	utils.SuccessResponse(c, list)
}
func Registration(c *gin.Context) {
	var req models.UserCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = repositories.Registration(c.Request.Context(), req.FirstName, req.LastName, "customer", req.Password, req.Email)

	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "user created")
}
func DeleteUser(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err := repositories.DeleteUser(c.Request.Context(), id)
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(200, gin.H{
	// 	"success": true,
	// })
	utils.SuccessResponse(c, "user deleted")
}
func GetUser(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	req, err := repositories.GetUser(c.Request.Context(), id)
	if utils.ErrorCheck(c, err) {

	}
	// c.JSON(200, gin.H{
	// 	"success": true,
	// 	"data":    req,
	// })
	utils.SuccessResponse(c, req)
}
func UpdateUser(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	// var req models.UserResponse
	var req models.UserCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = repositories.UpdateUser(c.Request.Context(), id, req)
	if utils.ErrorCheck(c, err) {
		return
	}
	// c.JSON(200, gin.H{
	// 	"success": true,
	// })
	utils.SuccessResponse(c, "user updated")
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func Login(c *gin.Context) {
	user, err := repositories.GetUserEmail(
		c,
		c.Query("email"))
	// err handle
	if utils.ErrorCheck(c, err) {
		return
	}

	if user.Password != c.Query("password") {
		// err handle for pass
		c.JSON(400, gin.H{
			"error": "wrong password",
		})
		return
	}
	Token := GenerateSecureToken(32)
	err = repositories.InsertToken(c, Token, user.Id)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}
	utils.SuccessResponse(c, Token)
}
func Logout(c *gin.Context) {
	if repositories.CheckIsTokenReal(c, c.Query("token")) {
		err := repositories.DeleteToken(c, c.Query("token"))
		if err != nil {
			utils.ErrorResponse(c, err)
			return
		}
		utils.SuccessResponse(c, "logout")
	} else {
		utils.SuccessResponse(c, "non token")
	}
}
func UserRoutes(rg *gin.RouterGroup) {
	rg.GET("/logout", Logout)
	rg.POST("/login", Login)
	rg.GET("/admin/users", UserList)
	rg.POST("/registration", Registration)
	rg.DELETE("/admin/users/delete/:id", DeleteUser)
	rg.GET("/admin/users/get/:id", GetUser)
	rg.PUT("/admin/users/update/:id", UpdateUser)
}
