package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gochkarovabagul-debug/practice/internal/controllers"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
	"github.com/gochkarovabagul-debug/practice/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	connStr := "host=localhost user=postgres password=1234 port=5432 sslmode=disable"
	utils.ConnectDB(connStr)
	defer utils.GetDB().Close(context.Background())
	r := gin.Default()
	r.Use(Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	rg := r.Group("/api")
	controllers.UserRoutes(rg)
	controllers.CategoryRoutes(rg)
	controllers.PharmacyRoutes(rg)
	controllers.PharmacyMedicinesRoutes(rg)
	controllers.OrderRoutes(rg)
	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path != "/api/login" && c.Request.URL.Path != "/api/registration" {
			auth := c.GetHeader("Authorization")
			token := strings.TrimPrefix(auth, "Bearer ")
			token = strings.TrimSpace(token)
			userId, err := repositories.GetUserIdByToken(c.Request.Context(), token)
			var expiresAt time.Time
			expiresAt, err = repositories.GetExpiresAtByToken(c, token)
			if expiresAt.Before(time.Now()) {
				return
			}
			log.Print(token, userId, err)
			if userId == 0 || err != nil {
				c.JSON(400, "token missing")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
