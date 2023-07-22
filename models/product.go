package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Url         string
	Description string
	Title       string
}

//User        User    `gorm:"foreignKey:UserRefer"`
//CompanyID   uint    `gorm:"foreignKey:CompanyID"`
//Company     Company `gorm:"foreignKey:CompanyRefer"`
//CompanyID   Company `gorm:"foreignKey:CompanyId"`
//User        User    `gorm:"foreignKey:UserId"`
