package models

import "gorm.io/gorm"

type AuthRequest struct {
	gorm.Model
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
