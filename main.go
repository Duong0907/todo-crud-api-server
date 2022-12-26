package main

import (
	"net/http"
	"strconv"
	"todo-crud-api-server/model"
	"todo-crud-api-server/database"
	"github.com/gin-gonic/gin"
)

func getTodos(c *gin.Context) {
	enableCors(c)
	todos := db.GetTodos()
	c.IndentedJSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
	enableCors(c)
	id, _ := strconv.Atoi(c.Param("id"))
	todo := db.GetTodo(id)
	c.IndentedJSON(http.StatusOK, todo)
}

func addTodo(c *gin.Context) {
	enableCors(c)
	var newTodo todo.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo error"})
		return
	}

	db.AddTodo(newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func removeTodo(c *gin.Context) {
	enableCors(c)
	id, _ := strconv.Atoi(c.Param("id"))
	db.RemoveTodo(id)
	todos := db.GetTodos()
	c.IndentedJSON(http.StatusOK, todos)
}

func updateTodo(c *gin.Context) {
	enableCors(c)
	id, _ := strconv.Atoi(c.Param("id"))

	var newTodo todo.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo error"})
	} else {
		db.UpdateTodo(id, newTodo)
		c.IndentedJSON(http.StatusOK, newTodo)
	}
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todo/:id", getTodo)
	router.POST("/todo", addTodo)
	router.DELETE("/todo/:id", removeTodo)
	router.PUT("/todo/:id", updateTodo)

	router.Run("10.0.158.232:8000")
}

func enableCors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
}
