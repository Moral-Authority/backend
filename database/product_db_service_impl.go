package database

import (
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
)

type ProductDbServiceImpl struct{}

func (s ProductDbServiceImpl) GetHomeGardenProductByID(productId string) (*models.HomeGardenProduct, error) {
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

func (s ProductDbServiceImpl) GetBathBeautyProductByID(productId string) (*models.HealthBathBeautyProduct, error) {
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

func (s ProductDbServiceImpl) GetClothingAccessoriesProductByID(productId string) (*models.ClothingAccessoriesProduct, error) {
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

func (s ProductDbServiceImpl) GetToysKidsBabiesProductByID(productId string) (*models.ToysKidsBabiesProduct, error) {
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

func (s ProductDbServiceImpl) GetAllHomeGardenProducts() ([]*models.HomeGardenProduct, error) {
	var products []*models.HomeGardenProduct
	result := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo").
		Find(&products)

	if result.Error != nil {
		logrus.Errorf("Unable to get all products, %s", result.Error)
		return nil, result.Error
	}

	logrus.Infof("Executed SQL: %s", result.Statement.SQL.String()) // Log the raw SQL query
	logrus.Infof("Length of prods: %d", len(products)) // Log the number of products found

	if len(products) == 0 {
		logrus.Warn("No products found")
	}

	return products, nil
}


func (s ProductDbServiceImpl) GetAllBathBeautyProducts() ([]*models.HealthBathBeautyProduct, error) {
	var products []*models.HealthBathBeautyProduct
	result := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo").
		Find(&products)
	if result.Error != nil {
		logrus.Errorf("Unable to get all products, %s", result.Error)
		return nil, result.Error
	}
	return products, nil
}

func (s ProductDbServiceImpl) GetAllClothingAccessoriesProducts() ([]*models.ClothingAccessoriesProduct, error) {
	var products []*models.ClothingAccessoriesProduct
	result := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo").
		Find(&products)
	if result.Error != nil {
		logrus.Errorf("Unable to get all products, %s", result.Error)
		return nil, result.Error
	}
	return products, nil
}

func (s ProductDbServiceImpl) GetAllToysKidsBabiesProducts() ([]*models.ToysKidsBabiesProduct, error) {
	var products []*models.ToysKidsBabiesProduct
	result := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo").
		Find(&products)
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
