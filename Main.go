package main

import (
	"github.com/gin-gonic/gin"
	"todo/database"
	"todo/handlers"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.GET("/todos", handlers.GetTodos)
	r.GET("/todos/:id", handlers.GetTodo)
	r.POST("/todos", handlers.CreateTodo)
	r.PUT("/todos/:id", handlers.UpdateTodo)

	r.DELETE("/todos/:id", handlers.DeleteTodo)

	r.Run(":8080")
}
