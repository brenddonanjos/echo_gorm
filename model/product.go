package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id          int       `json:"id"`
	BarCode     string    `json:"bar_code"`
	Name        string    `json:"name"`
	Value       float64   `json:"value"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	Active      int       `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Products []Product
