package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"practice/internal/controllers"
	"practice/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	//db connection
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s", "postgres", "practice_db", "postgres", "localhost", "5432")
	utils.ConnectDB(connStr)
	defer utils.GetDB().Close(context.Background())

	//HTTP serve
	r := gin.Default()

	rg := r.Group("/api")
	controllers.UserRoutes(rg)
	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
