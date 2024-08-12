package model

import "time"

// Product is a struct that represents a product
type Product struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Price         float64   `json:"price"`
	Description   string    `json:"description"`
	StockQuantity int       `json:"stock_quantity"`
	CategoryID    int       `json:"category_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
