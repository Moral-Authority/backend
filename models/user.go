package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	LoginCredentialsId uint
	LoginCredentials   LoginCredentials `gorm:"foreignKey:LoginCredentialsId"`
	Favorites         []Favorite      `gorm:"foreignKey:UserRefer"`
}

type LoginCredentials struct {
	gorm.Model
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Salt         string `json:"salt"`
}
