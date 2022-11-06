package models

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Title         string
	CategoryRefer uint
	Styles        []Style `gorm:"foreignKey:TypeRefer"`
}
