package main

import (
	"log"
	"os"

	"todo-crud-api-server/middlewares"
	"todo-crud-api-server/routes"
	"todo-crud-api-server/database"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("Fail to connect to database\nError: %v", err)
	}
	defer db.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	
	router := gin.Default()
	router.Use(middleware.CorsMiddleware)
	route.Setup(router)
	router.Run(":" + port)
}
