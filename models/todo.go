package models

type Todo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserID    uint   `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
