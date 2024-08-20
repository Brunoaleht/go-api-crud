package model

import "time"

type CarStatus string

const (
	CarStatusActive   CarStatus = "active"
	CarStatusInactive CarStatus = "inactive"
)

// Car is a struct that represents a car
type Car struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Status    CarStatus `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarWithProducts struct {
	ID        int          `json:"id"`
	UserID    int          `json:"user_id"`
	Status    CarStatus    `json:"status"`
	Products  []CarProduct `json:"products"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
