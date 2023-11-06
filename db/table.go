package snsdb

import (
	"time"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	Id        uint `gorm:"primaryKey"`
	Content   string
	UserId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
