package database

import "github.com/howstrongiam/backend/models"

type ProductDbService interface {
	GetProduct(productId string) *models.Product
	AddCategory(typeToAdd models.Categories) *models.Categories
	GetAllCategories() []*models.Categories
}
