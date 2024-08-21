package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title                 string
	Description           string
	Url                   string
	ProductCertifications []ProductCertification `gorm:"foreignKey:ProductID"`
	CompanyProducts       []CompanyProduct       `gorm:"foreignKey:ProductID"`
}
