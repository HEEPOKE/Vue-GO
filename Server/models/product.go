package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}
