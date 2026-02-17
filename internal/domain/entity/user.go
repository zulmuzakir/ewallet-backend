package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FullName  string
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(fullName, username, password, role string) *User {
	return &User{
		ID:        uuid.New(),
		FullName:  fullName,
		Username:  username,
		Password:  password,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
