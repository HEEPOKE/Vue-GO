package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Tel      string `json:"tel"`
	Role     int    `json:"role"`
}
