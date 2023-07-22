package models

import "gorm.io/gorm"

type Favourite struct {
	gorm.Model
	UserRefer uint
	ProductId uint `gorm:"foreignKey:ProductId"`
	Product   Product
}

//Product   Product `gorm:"foreignKey:ProductId"`
