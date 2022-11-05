package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name   string `json:"name"`
	TypeId uint   `json:"type_id"`
}
