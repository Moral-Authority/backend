package database

import (
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/volatiletech/null/v8"
)

type ProductDbService interface {
	AddProduct(product interface{}) (*interface{}, error)
	UpdateProduct(product model.UpdateProductRequest) (*interface{}, error)

	GetHomeGardenProductsByFilter(filters map[string]interface{}, subDepartment int) ([]*models.HomeGardenProduct, error)
	GetBathBeautyProductsByFilter(filters map[string]interface{}, subDepartment int) ([]*models.HealthBathBeautyProduct, error)
	GetClothingAccessoriesProductsByFilter(filters map[string]interface{}, subDepartment int) ([]*models.ClothingAccessoriesProduct, error)
	GetToysKidsBabiesProductsByFilter(filters map[string]interface{}, subDepartment int) ([]*models.ToysKidsBabiesProduct, error)

	GetHomeGardenProductByID(productId uint) (*models.HomeGardenProduct, error)
	GetHealthBathBeautyProductByID(productId uint) (*models.HealthBathBeautyProduct, error)
	GetClothingAccessoriesProductByID(productId uint) (*models.ClothingAccessoriesProduct, error)
	GetToysKidsBabiesProductByID(productId uint) (*models.ToysKidsBabiesProduct, error)

	// Updated functions with optional string filter
	GetRecentlyAddedProducts() ([]*models.HomeGardenProduct, error)
	GetAllHomeGardenProducts(filter null.Int) ([]*models.HomeGardenProduct, error)
	GetAllBathBeautyProducts(filter null.Int) ([]*models.HealthBathBeautyProduct, error)
	GetAllClothingAccessoriesProducts(filter null.Int) ([]*models.ClothingAccessoriesProduct, error)
	GetAllToysKidsBabiesProducts(filter null.Int) ([]*models.ToysKidsBabiesProduct, error)

	// Filters - Companies
	GetCompaniesFromHomeGarden(subDepartment int) ([]*string, error)
	GetCompaniesFromClothingAccessories(subDepartment int) ([]*string, error)
	GetCompaniesFromHealthBathBeauty(subDepartment int) ([]*string, error)
	GetCompaniesFromToysKidsBabies(subDepartment int) ([]*string, error)

	// Filters - Certifications - Companies
	GetCompanyCertificationsFromHomeGarden(subDepartment int) ([]*string, error)
	GetCompanyCertificationsFromHealthBathBeauty(subDepartment int) ([]*string, error)
	GetCompanyCertificationsFromClothingAccessories(subDepartment int) ([]*string, error)
	GetCompanyCertificationsFromToysKidsBabies(subDepartment int) ([]*string, error)

	// Filters - Certifications - Products
	GetProductCertificationsFromHomeGarden(subDepartment int) ([]*string, error)
	GetProductCertificationsFromHealthBathBeauty(subDepartment int) ([]*string, error)
	GetProductCertificationsFromClothingAccessories(subDepartment int) ([]*string, error)
	GetProductCertificationsFromToysKidsBabies(subDepartment int) ([]*string, error)

	// Filters - Price Range
	GetPriceRangeFromHomeGarden(subDepartment int) (*model.PriceRange, error)
	GetPriceRangeFromHealthBathBeauty(subDepartment int) (*model.PriceRange, error)
	GetPriceRangeFromClothingAccessories(subDepartment int) (*model.PriceRange, error)
	GetPriceRangeFromToysKidsBabies(subDepartment int) (*model.PriceRange, error)

	GetProductCompanyCertificationsFromHomeGarden(productID uint) ([]*models.Certification, error)
	GetProductCompanyCertificationsFromHealthBathBeauty(productID uint) ([]*models.Certification, error)
	GetProductCompanyCertificationsFromClothingAccessories(productID uint) ([]*models.Certification, error)
	GetProductCompanyCertificationsFromToysKidsBabies(productID uint) ([]*models.Certification, error)


	GetProductCertificationsFromHomeAndGardenByProduct(productID uint) ([]*models.Certification, error) 
	GetProductCertificationsFromHealthBathBeautyByProduct(productID uint) ([]*models.Certification, error) 
	GetProductCertificationsFromClothingAccessoriesByProduct(productID uint) ([]*models.Certification, error) 
	GetProductCertificationsFromToysKidsBabiesByProduct(productID uint) ([]*models.Certification, error) 

}

// 	AddProductCertification(productCert models.ProductCertification) (*models.ProductCertification, error)
// 	GetAllProductsByDepartment(department string, subDepartment null.String) ([]*interface{}, error)
// 	GetProductByID(productId string, department string) (*interface{}, error)
