package models

import (
	"errors"
	"time"
)

type User struct {
	ID           int        `db:"id" json:"id"`
	Username     string     `db:"username" json:"username"`
	Email        string     `db:"email" json:"email"`
	PasswordHash string     `db:"password_hash" json:"password_hash"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
}

func ValidateUser(user User) error {
	if user.Username == "" {
		return errors.New("username cannot be blank")
	}
	if len(user.Username) > 30 {
		return errors.New("username must be at most 30 characters")
	}
	if user.Email == "" {
		return errors.New("email cannot be blank")
	}
	if len(user.Email) > 150 {
		return errors.New("email cannot be more than 150 characters")
	}
	return nil
}
