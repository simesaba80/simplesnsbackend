package db

import (
	"time"
)

type User struct { //コメントはDB内のカラム名
	Id        uint      `gorm:"primaryKey"` //id
	UserID    string    //user_id
	Name      string    //name
	Email     string    //email
	Password  string    //password
	CreatedAt time.Time //created_at
	UpdatedAt time.Time //updated_at
}

type Post struct { //コメントはDB内のカラム名
	Id        uint      `gorm:"primaryKey"` //id
	Content   string    //content
	UserId    uint      //user_id
	CreatedAt time.Time //created_at
	UpdatedAt time.Time //updated_at
}

type Friend struct {
	Id       uint `gorm:"primaryKey"`
	UserID   uint
	FriendID uint
}
