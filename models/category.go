package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	DepartmentId uint   `json:"department_id"`
	Title        string `json:"title"`
}
