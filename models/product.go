package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Url             string
	Description     string
	UserId          uint
	User            User `gorm:"foreignKey:UserId"`
	ImageId         uint
	Image           Image `gorm:"foreignKey:ImageId"`
	CertificationId uint
	Certification   Certification `gorm:"foreignKey:CertificationId"`
	StyleRefer      uint
}
