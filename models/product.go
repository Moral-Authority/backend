package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Url         string
	Description string
	Title       string
	Company     Company
	User        User
	//CompanyID   Company `gorm:"foreignKey:CompanyId"`
	//User        User    `gorm:"foreignKey:UserId"`
}
