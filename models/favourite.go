package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserRefer uint
	User      User   `gorm:"foreignKey:UserRefer"`
	ProductId uint
	Product   Product `gorm:"foreignKey:ProductId"`
}