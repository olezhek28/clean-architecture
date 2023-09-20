package model

import (
	"database/sql"
	"time"
)

type User struct {
	UUID      string
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserInfo struct {
	FirstName string
	LastName  string
	Age       int64
}
