package handlers

import (
	"net/http"
	"todo/database"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)

	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Todo not Found",
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	database.DB.Create(&todo)

	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Todo not found",
		})
		return
	}

	var input models.Todo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	todo.Title = input.Title
	todo.Completed = input.Completed

	database.DB.Save(&todo)

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Todo not found",
		})
		return
	}

	database.DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted",
	})
}
