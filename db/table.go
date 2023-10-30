package snsdb

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type Post struct {
	gorm.Model
	Content string
	UserId  uint
}
