package models

import (
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

type CompanyProduct struct {
	gorm.Model
	CompanyID      uint        `gorm:"foreignKey:CompanyRefer"`
	ProductID      uint        `gorm:"foreignKey:ProductRefer"`
	LaunchedAt     null.Time   `json:"launched_at"`
	DiscontinuedAt null.Time   `json:"discontinued_at"`
	OtherDetails   null.String `json:"other_details"`
}
