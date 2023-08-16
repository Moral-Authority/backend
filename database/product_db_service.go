package database

import "github.com/Moral-Authority/backend/models"

type ProductDbService interface {
	GetProduct(productId string) *models.Product
	AddCategory(typeToAdd models.Category) *models.Category
	GetAllCategories() []*models.Category
}
