package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Title      string     `json:"title"`
	Categories []Category `gorm:"foreignKey:DepartmentRefer"`
}
