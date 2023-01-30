package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Title      string     `json:"title"`
	Categories []Category `gorm:"foreignKey:DepartmentRefer"`
}

// TODO check foreign keys here
type Section struct {
	gorm.Model
	Title       string       `json:"title"`
	SubSections []SubSection `gorm:"foreignKey:SubSectionRefer"`
}

type SubSection struct {
	gorm.Model
	Title       string       `json:"title"`
	Departments []Department `gorm:"foreignKey:DepartmentRefer"`
}
