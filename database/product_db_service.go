package database

import (
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
)

type ProductDbService interface {
	GetProductByID(productId string) (*models.Product, error)
	GetAllProducts() ([]*models.Product, error)
	AddProduct(product models.Product) (*models.Product, error)
	DeleteProduct(productId string) error
	UpdateProduct(product model.UpdateProductRequest) (*models.Product, error)
	GetProductsByFilter(filters map[string]interface{}) ([]models.Product, error)
	AddCategory(category models.Category) (*models.Category, error)
	GetAllCategories() ([]*models.Category, error)
	AddProductCertification(productCert models.ProductCertification) (*models.ProductCertification, error)
	
}
