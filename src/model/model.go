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
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
