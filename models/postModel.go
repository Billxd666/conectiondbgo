package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    *string
	Password string
}

type Post struct {
	gorm.Model
	Id        uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
