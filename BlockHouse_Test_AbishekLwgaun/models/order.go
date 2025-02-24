package models

import "gorm.io/gorm"

// just the defining the db schema
type Order struct {
	gorm.Model
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	OrderType string  `json:"order_type"`
}
