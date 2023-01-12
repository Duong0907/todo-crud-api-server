package main

import (
	"log"

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
	
	router := gin.Default()
	router.Use(middleware.CorsMiddleware)
	route.Setup(router)
	router.Run(":8000")
}

// connect command in mongosh:
// mongosh "mongodb+srv://cluster0.aszqp6w.mongodb.net/myFirstDatabase" --apiVersion 1 --username Duong
