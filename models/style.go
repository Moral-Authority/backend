package models

import "gorm.io/gorm"

type Style struct {
	gorm.Model
	Title     string
	TypeRefer uint
	Products  []Product `gorm:"foreignKey:StyleRefer"`
}
