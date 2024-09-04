package database

import (
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/volatiletech/null/v8"
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
	// Updated functions with optional string filter
	GetAllHomeGardenProducts(filter null.Int) ([]*models.HomeGardenProduct, error)
	GetAllBathBeautyProducts(filter null.Int) ([]*models.HealthBathBeautyProduct, error)
	GetAllClothingAccessoriesProducts(filter null.Int) ([]*models.ClothingAccessoriesProduct, error)
	GetAllToysKidsBabiesProducts(filter null.Int) ([]*models.ToysKidsBabiesProduct, error)
}

// 	AddProductCertification(productCert models.ProductCertification) (*models.ProductCertification, error)
// 	GetAllProductsByDepartment(department string, subDepartment null.String) ([]*interface{}, error)
// 	GetProductByID(productId string, department string) (*interface{}, error)
