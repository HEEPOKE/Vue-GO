package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code        int64  `json:"code" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}
