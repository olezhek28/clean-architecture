package model

import (
	"time"
)

type User struct {
	UUID      string
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UserInfo struct {
	FirstName string
	LastName  string
	Age       int64
}
