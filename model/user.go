package model

import "time"

// User is a struct that represents a user
type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	GroupID      int       `json:"group_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestUpdatePassword struct {
	Email string `json:"email"`
}

type UpdatePassword struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}
