package database

import (
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
)

type ProductDbService interface {
	AddProduct(product interface{}) (*interface{}, error)
	UpdateProduct(product model.UpdateProductRequest) (*interface{}, error)
	GetHomeGardenProductsByFilter(filters map[string]interface{}) ([]*models.HomeGardenProduct, error)
	GetBathBeautyProductsByFilter(filters map[string]interface{}) ([]*models.HealthBathBeautyProduct, error)
	GetClothingAccessoriesProductsByFilter(filters map[string]interface{}) ([]*models.ClothingAccessoriesProduct, error)
	GetToysKidsBabiesProductsByFilter(filters map[string]interface{}) ([]*models.ToysKidsBabiesProduct, error)
	GetHomeGardenProductByID(productId uint) (*models.HomeGardenProduct, error)
	GetHealthBathBeautyProductByID(productId uint) (*models.HealthBathBeautyProduct, error)
	GetClothingAccessoriesProductByID(productId uint) (*models.ClothingAccessoriesProduct, error)
	GetToysKidsBabiesProductByID(productId uint) (*models.ToysKidsBabiesProduct, error)
	GetAllHomeGardenProducts() ([]*models.HomeGardenProduct, error)
	GetAllBathBeautyProducts() ([]*models.HealthBathBeautyProduct, error)
	GetAllClothingAccessoriesProducts() ([]*models.ClothingAccessoriesProduct, error)
	GetAllToysKidsBabiesProducts() ([]*models.ToysKidsBabiesProduct, error)
}

// 	AddProductCertification(productCert models.ProductCertification) (*models.ProductCertification, error)
// 	GetAllProductsByDepartment(department string, subDepartment null.String) ([]*interface{}, error)
// 	GetProductByID(productId string, department string) (*interface{}, error)
