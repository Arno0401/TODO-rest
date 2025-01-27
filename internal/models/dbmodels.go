package models

import "time"

type Todo struct {
	ID          int       `gorm:"id"`
	UserID      int       `gorm:"user_id"`
	Title       string    `gorm:"title"`
	Description string    `gorm:"description"`
	Status      bool      `gorm:"status"`
	CreatedAt   time.Time `gorm:"created_at"`
}

type Users struct {
	ID       int    `gorm:"id"`
	UserName string `gorm:"user_name"`
	Login    string `gorm:"login"`
	Password string `gorm:"password"`
	Role     string `gorm:"default:user"`
}
