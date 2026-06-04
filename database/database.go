package database

import (
	"todo/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Todo{})

	DB = db
}
