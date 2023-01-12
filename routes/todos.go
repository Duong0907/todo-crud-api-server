package route

import (
	"github.com/gin-gonic/gin"
	"todo-crud-api-server/controllers"
)

func Setup(router *gin.Engine) {
	router.GET("/api/todos", controller.GetTodos)
	router.GET("/api/todo/:id", controller.GetTodo)
	router.POST("/api/todo", controller.AddTodo)
	router.DELETE("/api/todo/:id", controller.RemoveTodo)
	router.PUT("/api/todo/:id", controller.UpdateTodo)
}
