package models

import (
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name        string      `gorm:"column:name json:name"`
	Url         null.String `gorm:"column:url" json:"url"`
	Description null.String `gorm:"column:description" json:"description"`
	UserId      null.Int64  `gorm:"column:user_id" json:"user_id"`
	IsVerified  null.Bool   `gorm:"column:is_verified" json:"is_verified"`
	City        null.String `gorm:"column:city" json:"city"`
	State       null.String `gorm:"column:state" json:"state"`
	Country     null.String `gorm:"column:country" json:"country"`
	ImageId     null.Int64  `gorm:"column:image_id" json:"image_id"`
	Image       null.String `gorm:"foreignKey:ImageId"`
}

type Company_Certs struct {
	gorm.Model
	CompanyID       uint    `gorm:"foreignKey:CompanyID"`
	Company         Company `gorm:"foreignKey:CompanyRefer"`
	CertificationId Certification
	//CompanyId       Company
	//CertificationId Certification
}
