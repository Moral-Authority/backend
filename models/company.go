package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Url             string
	Description     string
	UserId          uint
	User            User `gorm:"foreignKey:UserId"`
	IsVerified      bool
	ImageId         uint
	Image           Image `gorm:"foreignKey:ImageId"`
	CertificationId uint
	Certification   Certification `gorm:"foreignKey:CertificationId"`
}
