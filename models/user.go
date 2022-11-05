package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	LoginCredentialsId uint
	LoginCredentials   LoginCredentials `gorm:"foreignKey:LoginCredentialsId"`
	Favourites         []Favourite      `gorm:"foreignKey:UserRefer"`
}

type LoginCredentials struct {
	gorm.Model
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Salt         string `json:"salt"`
}
