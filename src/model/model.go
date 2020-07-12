package model

import (
	"time"
)

type User struct {
	ID        int32
	Name      string
	Email     string
	PhotoUrl  string
	Password  string
	Comments []Comment `gorm:"foreignkey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Comment struct{
	ID        int32
	UserId    int32
	PostId    int32
	Content   string `gorm:"size:100"`
}
