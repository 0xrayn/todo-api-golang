package handlers

import (
	"net/http"
	"todo/database"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	// Ambil userID dari token
	userID := c.MustGet("userID")

	var todos []models.Todo
	database.DB.Where("user_id = ?", userID).Find(&todos)

	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	userID := c.MustGet("userID")
	id := c.Param("id")

	var todo models.Todo
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Todo tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	// Ambil userID dari token
	userID := c.MustGet("userID")

	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	todo := models.Todo{
		UserID:    uint(userID.(float64)),
		Title:     input.Title,
		Completed: input.Completed,
	}

	database.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	userID := c.MustGet("userID")
	id := c.Param("id")

	var todo models.Todo
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Todo tidak ditemukan",
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
	userID := c.MustGet("userID")
	id := c.Param("id")

	var todo models.Todo
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Todo tidak ditemukan",
		})
		return
	}

	database.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo berhasil dihapus",
	})
}
