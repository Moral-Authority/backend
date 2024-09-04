package database

import (
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
)

type ProductDbServiceImpl struct{}

func (s ProductDbServiceImpl) GetHomeGardenProductByID(productId uint) (*models.HomeGardenProduct, error) {
	var product models.HomeGardenProduct
	result := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo").
		First(&product, "id = ?", productId)
	if result.Error != nil {
		logrus.Errorf("Unable to get HomeGarden product, %s", result.Error)
		return nil, result.Error
	}
	return &product, nil
}

func (s ProductDbServiceImpl) GetHealthBathBeautyProductByID(productId uint) (*models.HealthBathBeautyProduct, error) {
	var product models.HealthBathBeautyProduct
	result := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo").
		First(&product, "id = ?", productId)
	if result.Error != nil {
		logrus.Errorf("Unable to get BathBeauty product, %s", result.Error)
		return nil, result.Error
	}
	return &product, nil
}

func (s ProductDbServiceImpl) GetClothingAccessoriesProductByID(productId uint) (*models.ClothingAccessoriesProduct, error) {
	var product models.ClothingAccessoriesProduct
	result := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo").
		First(&product, "id = ?", productId)
	if result.Error != nil {
		logrus.Errorf("Unable to get ClothingAccessories product, %s", result.Error)
		return nil, result.Error
	}
	return &product, nil
}

func (s ProductDbServiceImpl) GetToysKidsBabiesProductByID(productId uint) (*models.ToysKidsBabiesProduct, error) {
	var product models.ToysKidsBabiesProduct
	result := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo").
		First(&product, "id = ?", productId)
	if result.Error != nil {
		logrus.Errorf("Unable to get ToysKidsBabies product, %s", result.Error)
		return nil, result.Error
	}
	return &product, nil
}


// GetAllHomeGardenProducts retrieves all HomeGarden products with an optional subDepartment filter
func (s ProductDbServiceImpl) GetAllHomeGardenProducts(subDepartment null.Int) ([]*models.HomeGardenProduct, error) {
	var products []*models.HomeGardenProduct
	query := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo")

	// Apply the subDepartment filter if it's provided and valid
	if subDepartment.Valid {
		query = query.Where("sub_department = ?", subDepartment.Int)
	}

	result := query.Find(&products)
	if result.Error != nil {
		logrus.Errorf("Unable to get all products, %s", result.Error)
		return nil, result.Error
	}

	if len(products) == 0 {
		logrus.Warn("No products found")
	}

	return products, nil
}

// GetAllBathBeautyProducts retrieves all HealthBathBeauty products with an optional subDepartment filter
func (s ProductDbServiceImpl) GetAllBathBeautyProducts(subDepartment null.Int) ([]*models.HealthBathBeautyProduct, error) {
	var products []*models.HealthBathBeautyProduct
	query := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo")

	// Apply the subDepartment filter if it's provided and valid
	if subDepartment.Valid {
		query = query.Where("sub_department = ?", subDepartment.Int)
	}

	result := query.Find(&products)
	if result.Error != nil {
		logrus.Errorf("Unable to get all products, %s", result.Error)
		return nil, result.Error
	}

	return products, nil
}

// GetAllClothingAccessoriesProducts retrieves all ClothingAccessories products with an optional subDepartment filter
func (s ProductDbServiceImpl) GetAllClothingAccessoriesProducts(subDepartment null.Int) ([]*models.ClothingAccessoriesProduct, error) {
	var products []*models.ClothingAccessoriesProduct
	query := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo")

	// Apply the subDepartment filter if it's provided and valid
	if subDepartment.Valid {
		query = query.Where("sub_department = ?", subDepartment.Int)
	}

	result := query.Find(&products)
	if result.Error != nil {
		logrus.Errorf("Unable to get all products, %s", result.Error)
		return nil, result.Error
	}

	return products, nil
}

// GetAllToysKidsBabiesProducts retrieves all ToysKidsBabies products with an optional subDepartment filter
func (s ProductDbServiceImpl) GetAllToysKidsBabiesProducts(subDepartment null.Int) ([]*models.ToysKidsBabiesProduct, error) {
	var products []*models.ToysKidsBabiesProduct
	query := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo")

	// Apply the subDepartment filter if it's provided and valid
	if subDepartment.Valid {
		query = query.Where("sub_department = ?", subDepartment.Int)
	}

	result := query.Find(&products)
	if result.Error != nil {
		logrus.Errorf("Unable to get all products, %s", result.Error)
		return nil, result.Error
	}

	return products, nil
}


func (s ProductDbServiceImpl) GetHomeGardenProductsByFilter(filters map[string]interface{}) ([]*models.HomeGardenProduct, error) {
	var products []*models.HomeGardenProduct
	db := GetDbConn()

	query := ApplyFilters(db, filters)

	if err := query.Find(&products).Error; err != nil {
		logrus.Errorf("Unable to get HomeGarden products by filter, %s", err)
		return nil, err
	}

	return products, nil
}

func (s ProductDbServiceImpl) GetBathBeautyProductsByFilter(filters map[string]interface{}) ([]*models.HealthBathBeautyProduct, error) {
	var products []*models.HealthBathBeautyProduct
	db := GetDbConn()

	query := ApplyFilters(db, filters)

	if err := query.Find(&products).Error; err != nil {
		logrus.Errorf("Unable to get BathBeauty products by filter, %s", err)
		return nil, err
	}

	return products, nil
}

func (s ProductDbServiceImpl) GetClothingAccessoriesProductsByFilter(filters map[string]interface{}) ([]*models.ClothingAccessoriesProduct, error) {
	var products []*models.ClothingAccessoriesProduct
	db := GetDbConn()

	query := ApplyFilters(db, filters)

	if err := query.Find(&products).Error; err != nil {
		logrus.Errorf("Unable to get ClothingAccessories products by filter, %s", err)
		return nil, err
	}

	return products, nil
}

func (s ProductDbServiceImpl) GetToysKidsBabiesProductsByFilter(filters map[string]interface{}) ([]*models.ToysKidsBabiesProduct, error) {
	var products []*models.ToysKidsBabiesProduct
	db := GetDbConn()

	query := ApplyFilters(db, filters)

	if err := query.Find(&products).Error; err != nil {
		logrus.Errorf("Unable to get ToysKidsBabies products by filter, %s", err)
		return nil, err
	}

	return products, nil
}


func (s ProductDbServiceImpl) UpdateProduct(product model.UpdateProductRequest) (*interface{}, error) {

	return nil, nil
}

func (s ProductDbServiceImpl) AddProduct(product interface{}) (*interface{}, error) {
	var err error

	err = GetDbConn().Create(product).Error
	if err != nil {
		logrus.Errorf("Unable to save product, %s", err)
		return nil, err
	}

	return &product, nil
}
