package models

import "gorm.io/gorm"

type SectionType struct {
	gorm.Model
	SectionId uint `json:"section_id"`
	TypeId    uint `json:"type_id"`
}
