package main

import (
	"github.com/gin-gonic/gin"
	"todo/database"
	"todo/handlers"
	"todo/middleware"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	{

		auth.GET("/todos", handlers.GetTodos)
		auth.GET("/todos/:id", handlers.GetTodo)
		auth.POST("/todos", handlers.CreateTodo)
		auth.PUT("/todos/:id", handlers.UpdateTodo)

		auth.DELETE("/todos/:id", handlers.DeleteTodo)
	}

	r.Run(":8080")
}
