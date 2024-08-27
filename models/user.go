package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string      `json:"email" gorm:"unique"` 
	Phone        *string      `json:"phone" gorm:"unique"` 
	PasswordHash string      `json:"password_hash"`
	Salt         string      `json:"salt"`
	Favorites    []Favorite  `gorm:"foreignKey:UserRefer"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
