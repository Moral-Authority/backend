package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	DepartmentRefer uint
	Title           string
	Types           []Type `gorm:"foreignKey:CategoryRefer"`
}

type SubCategory struct {
	gorm.Model
	CategoryRefer uint
	Title         string
	Departments   []Type `gorm:"foreignKey:CategoryRefer"`
}
