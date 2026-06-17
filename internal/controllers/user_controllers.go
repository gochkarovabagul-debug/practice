package controllers

import (
	"strings"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/services"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func UserList(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	role := c.Query("role")
	list, err := services.UserListService(c, models.UserFilter{
		Limit:  limit,
		Offset: offset,
		Search: search,
		Role:   role,
	})

	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, list)
}
func Registration(c *gin.Context) {
	var req models.UserCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = services.RegistrationService(c, req.FirstName, req.LastName, "customer", req.Password, req.Email)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "user created")
}
func DeleteUser(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	err := services.DeleteUserService(c, id)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "user deleted")
}
func GetUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.TrimPrefix(auth, "Bearer ")
	token = strings.TrimSpace(token)
	req, err := services.GetUserService(c, token, false)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, req)
}
func UpdateUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.TrimPrefix(auth, "Bearer ")
	token = strings.TrimSpace(token)
	var req models.UserUpdateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = services.UpdateUserService(c, token, req)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, "user updated")
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	err := c.Bind(&req)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}
	Token, err := services.LoginService(c, req.Email, req.Password)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}
	user, err := services.GetUserService(c, Token, false)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}
	utils.SuccessResponse(c, gin.H{
		"token":      Token,
		"expires_at": time.Now().AddDate(1, 0, 0), // TODO: change this
		"user":       user,
	})
}
func Logout(c *gin.Context) {
	// token := c.Query("token")
	auth := c.GetHeader("Authorization")
	token := strings.TrimPrefix(auth, "Bearer ")
	token = strings.TrimSpace(token)
	message, err := services.LogoutService(c, token)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}
	utils.SuccessResponse(c, message)
}
func ChangePassword(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.TrimPrefix(auth, "Bearer ")
	token = strings.TrimSpace(token)
	var passchange models.ChangePasswordRequest
	err := c.Bind(&passchange)
	if utils.ErrorCheck(c, err) {
		return
	}
	var req models.UserResponse
	err = services.ChangePasswordService(c, token, true, passchange, req)
	if utils.ErrorCheck(c, err) {
		utils.ErrorResponse(c, err)
		return
	}
	utils.SuccessResponse(c, "password changed")
}
func UserRoutes(rg *gin.RouterGroup) {
	rg.GET("/auth/logout", Logout)
	rg.POST("/auth/login", Login)
	rg.GET("/admin/users", UserList)
	rg.POST("/registration", Registration)
	rg.DELETE("/admin/users/delete/:id", DeleteUser)
	rg.GET("/user/me", GetUser)
	rg.POST("/user/me", UpdateUser)
	rg.POST("/user/changepassword", ChangePassword)
}
