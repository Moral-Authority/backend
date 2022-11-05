package models

import "gorm.io/gorm"

type ServiceType struct {
	gorm.Model
	TypeId    uint `json:"type_id"`
	ServiceId uint `json:"service_id"`
}
