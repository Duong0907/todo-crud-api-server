package controller

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-crud-api-server/models"
	"todo-crud-api-server/database"
)

func GetTodos(c *gin.Context) {
	todos, err := db.GetTodos()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := db.GetTodo(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func AddTodo(c *gin.Context) {
	var newTodo todo.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo error"})
		return
	}
	if err := db.AddTodo(newTodo); err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusCreated, newTodo)

}

func RemoveTodo(c *gin.Context) {
	id := c.Param("id")
	db.RemoveTodo(id)
	todos, err := db.GetTodos()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var newTodo todo.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo error"})
		return
	}
	if err := db.UpdateTodo(id, newTodo); err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, newTodo)
}