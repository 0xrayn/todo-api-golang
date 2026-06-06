package main

import (
	"github.com/gin-gonic/gin"
	"todo/database"
	"todo/handlers"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.GET("/todos", handlers.GetTodos)
	r.GET("/todos/:id", handlers.GetTodo)
	r.POST("/todos", handlers.CreateTodo)
	r.PUT("/todos/:id", handlers.UpdateTodo)

	r.DELETE("/todos/:id", handlers.DeleteTodo)

	r.Run(":8080")
}
