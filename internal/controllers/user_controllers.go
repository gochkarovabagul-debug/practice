package controllers

import (
	"strings"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/permission"
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
	list, total, err := services.UserListService(c, models.UserFilter{
		Limit:  limit,
		Offset: offset,
		Search: search,
		Role:   role,
	})

	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponseList(c, list, total, limit, offset)
}
func Registration(c *gin.Context) {
	var req models.UserCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	var token string
	token, err = services.RegistrationService(c, req.Name, "customer", req.Password, req.Email)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, token)
}
func CreateUserByAdmin(c *gin.Context) {
	var req models.UserCreateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	user, err := services.CreateUserByAdminService(c, req.Name, req.Role, req.Password, req.Email)
	utils.SuccessResponse(c, user)
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
func GetUserById(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	user, err := services.GetUserByIdService(c, id, false)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, user)
}
func GetUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.TrimPrefix(auth, "Bearer ")
	token = strings.TrimSpace(token)
	req, err := services.GetUserByTokenService(c, token, false)
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
func UpdateUserByAdmin(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	var req models.UserUpdateRequest
	err := c.Bind(&req)
	if utils.ErrorCheck(c, err) {
		return
	}
	err = services.UpdateUserByIdService(c, id, req)
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
	user, err := services.GetUserByTokenService(c, Token, false)
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
	gg := rg.Group("").Use(permission.RequireAdmin())
	rg.POST("/auth/register", Registration)
	rg.POST("/auth/login", Login)
	rg.POST("/auth/logout", Logout)
	gg.GET("/admin/users", UserList)
	gg.POST("/admin/users", CreateUserByAdmin)
	gg.GET("/admin/users/:id", GetUserById)
	gg.PUT("/admin/users/:id", UpdateUserByAdmin)
	gg.DELETE("/admin/users/:id", DeleteUser)
	rg.GET("/me", GetUser)
	rg.PUT("/me", UpdateUser)
	rg.POST("/user/changepassword", ChangePassword)
}
