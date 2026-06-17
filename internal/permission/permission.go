package permission

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

type Role string

const AdminRole Role = "admin"
const PharmacyRole Role = "pharmacy"
const UserRole Role = "user"

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		token := strings.TrimPrefix(auth, "Bearer ")
		token = strings.TrimSpace(token)
		fmt.Print(1)
		role, err := repositories.GetRoleByToken(c, token)
		if err != nil {
			utils.ErrorResponse(c, err)
		}
		fmt.Print(2)
		if role != string(AdminRole) {
			utils.ErrorResponse(c, errors.New("not admin"))
			c.Abort()
			return
		}
		c.Next()
	}
}
func RequirePharmacyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		token := strings.TrimPrefix(auth, "Bearer ")
		token = strings.TrimSpace(token)
		fmt.Print(3)
		role, err := repositories.GetRoleByToken(c, token)
		if err != nil {
			utils.ErrorResponse(c, err)
		}
		fmt.Print(4)
		if role != string(AdminRole) && role != string(PharmacyRole) {
			utils.ErrorResponse(c, errors.New("not admin"))
			c.Abort()
			return
		}
		c.Next()
		fmt.Print(5)
	}
}
