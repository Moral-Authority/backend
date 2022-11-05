package models

import "gorm.io/gorm"

type LoginCredentials struct {
	gorm.Model
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Salt         string `json:"salt"`
}
