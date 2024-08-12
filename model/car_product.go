package model

// CarProduct is a struct that represents a car product
type CarProduct struct {
	ID        int     `json:"id"`
	CarID     int     `json:"car_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}
