package models

import "gorm.io/gorm"

// type Favorite struct {
// 	gorm.Model
// 	UserRefer uint
// 	User      User   `gorm:"foreignKey:UserRefer"`
// 	ProductId uint
// 	Product   Product `gorm:"foreignKey:ProductId"`
// }

type Favorite struct {
	gorm.Model
	UserRefer         uint `gorm:"index"`
	ProductID         uint `gorm:"index"`
	ProductDepartment int  `gorm:"index"` // The type of product (e.g., "HomeGardenProduct")
}
