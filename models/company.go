package models

import (
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Url         null.String `gorm:"column:url" json:"url"`
	Description null.String `gorm:"column:description" json:"description"`
	UserId      null.Int64  `gorm:"column:user_id" json:"user_id"`
	IsVerified  null.Bool   `gorm:"column:is_verified" json:"is_verified"`
	ImageId     null.Int64  `gorm:"column:image_id" json:"image_id"`
	Image       null.String `gorm:"foreignKey:ImageId"`
}
