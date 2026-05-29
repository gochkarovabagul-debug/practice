package main

import (
	"context"
	"log"
	"os"

	"github.com/gochkarovabagul-debug/practice/internal/controllers"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
	"github.com/gochkarovabagul-debug/practice/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	//db connection
	// connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s", "postgres", "practice_db", "postgres", "localhost", "5432")
	connStr := "host=localhost user=postgres password=1234 port=5432 sslmode=disable"
	utils.ConnectDB(connStr)
	defer utils.GetDB().Close(context.Background())

	//HTTP serve
	r := gin.Default()
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
		if c.Request.URL.Path != "/api/login" && c.Request.URL.Path != "/api/registration" {
			token := c.Query("token")
			userId, err := repositories.GetUserIdByToken(c.Request.Context(), token)
			if userId == 0 || err != nil {
				c.JSON(400, "token missing")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
