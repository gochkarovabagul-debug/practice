package utils

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func ConnectDB(config string) {
	conn, err := pgxpool.New(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Db error : %v\n", err)
		os.Exit(1)
	}
	db = conn
}
func ErrorCheck(c *gin.Context, err error) bool {
	if err != nil {
		ErrorResponse(c, err)
		return true
	}
	return false
}
func ErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success":   false,
		"error_msg": err.Error(),
	})
}
func SuccessResponseList(c *gin.Context, message any, total int, limit int, offset int) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    message,
		"meta": gin.H{
			"total":  total,
			"limit":  limit,
			"offset": offset,
		},
	})
}
func SuccessResponse(c *gin.Context, message any) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    message,
	})
}

func GetDB() *pgxpool.Pool {
	return db
}
