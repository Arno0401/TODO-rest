package repository

import "gorm.io/gorm"

type Repository struct {
	conn *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{
		conn: conn,
	}
}
