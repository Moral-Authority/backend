package models

import "gorm.io/gorm"

type CompanyImage struct {
	gorm.Model
	CompanyId     uint   `json:"company_id"`
	ImageLocation string `json:"image_location"`
}
