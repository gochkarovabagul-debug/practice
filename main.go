package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gochkarovabagul-debug/practice/config"
	"github.com/gochkarovabagul-debug/practice/internal/controllers"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
	"github.com/joho/godotenv"
)

func main() {
	// connStr := "host=localhost user=postgres password=1234 port=5432 sslmode=disable"
	// utils.ConnectDB(connStr)
	config.LoadConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secretconnectionText := os.Getenv("DATABASE_URL")
	utils.ConnectDB(secretconnectionText)
	// defer utils.GetDB().Close()
	defer utils.GetDB().Close()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(Logger())
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

		fmt.Println(1)

		if c.Request.URL.Path != "/api/auth/login" && c.Request.URL.Path != "/api/registration" {
			fmt.Println(1)
			auth := c.GetHeader("Authorization")
			token := strings.TrimPrefix(auth, "Bearer ")
			token = strings.TrimSpace(token)
			fmt.Printf("token:%v", token)

			userId, err := repositories.GetUserIdByToken(c.Request.Context(), token)
			var expiresAt time.Time
			expiresAt, err = repositories.GetExpiresAtByToken(c, token)
			if expiresAt.Before(time.Now()) {
				c.JSON(400, "token missing")
				c.Abort()
				return
			}
			fmt.Println(1)
			log.Println(token, userId, err)
			if userId == 0 || err != nil {
				c.JSON(400, "token missing")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
