package database

import (
	"github.com/Moral-Authority/backend/models"
)

type ProductDbService interface {
	GetProduct(productId string) (*models.Product, error)
	GetAllProducts() ([]*models.Product, error)
	AddProduct(product models.Product) (*models.Product, error)
	DeleteProduct(productId string) error
	AddCategory(category models.Category) (*models.Category, error)
	GetAllCategories() ([]*models.Category, error)
}
