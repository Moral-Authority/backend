package database

import (
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
)

type ProductDbServiceImpl struct{}

func (s ProductDbServiceImpl) GetRecentlyAddedProducts() ([]*models.HomeGardenProduct, error) {
	var products []*models.HomeGardenProduct
	result := GetDbConn().
		Preload("Company").
		Preload("PurchaseInfo").
		Order("created_at desc").
		Limit(10).
		Find(&products)
	if result.Error != nil {
		logrus.Errorf("Unable to get recently added products, %s", result.Error)
		return nil, result.Error
	}
	
	return products, nil
}

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

func (s ProductDbServiceImpl) GetHomeGardenProductsByFilter(filters map[string]interface{}, subDept int) ([]*models.HomeGardenProduct, error) {
	var products []*models.HomeGardenProduct

	// Start building the query, with Preloads for related tables
	query := GetDbConn().Debug().
		Preload("Company").     // Preload the Company relationship
		Preload("PurchaseInfo") // Preload the PurchaseInfo relationship

	// Apply dynamic product filters
	query = ApplyProductFilters(query, filters, subDept, "home_garden_products")

	logrus.Infof("Query: %v", query)
	// Execute the query
	if err := query.Find(&products).Error; err != nil {
		logrus.Errorf("Unable to get HomeGarden products by filter, %s", err)
		return nil, err
	}

	logrus.Infof("Products: %v", len(products))

	return products, nil
}

func (s ProductDbServiceImpl) GetBathBeautyProductsByFilter(filters map[string]interface{}, subDept int) ([]*models.HealthBathBeautyProduct, error) {
	var products []*models.HealthBathBeautyProduct

	// Start building the query, with Preloads for related tables
	query := GetDbConn().Debug().
		Preload("Company").     // Preload the Company relationship
		Preload("PurchaseInfo") // Preload the PurchaseInfo relationship

	// Apply dynamic product filters
	query = ApplyProductFilters(query, filters,subDept, "health_bath_beauty_products")

	if err := query.Find(&products).Error; err != nil {
		logrus.Errorf("Unable to get HealthBathBeauty products by filter, %s", err)
		return nil, err
	}

	return products, nil
}

func (s ProductDbServiceImpl) GetClothingAccessoriesProductsByFilter(filters map[string]interface{}, subDept int) ([]*models.ClothingAccessoriesProduct, error) {
	var products []*models.ClothingAccessoriesProduct

	// Start building the query, with Preloads for related tables
	query := GetDbConn().Debug().
		Preload("Company").     // Preload the Company relationship
		Preload("PurchaseInfo") // Preload the PurchaseInfo relationship

	// Apply dynamic product filters
	query = ApplyProductFilters(query, filters, subDept,"clothing_accessories_products")

	// Execute the query
	if err := query.Find(&products).Error; err != nil {
		logrus.Errorf("Unable to get ClothingAccessories products by filter, %s", err)
		return nil, err
	}

	return products, nil
}

func (s ProductDbServiceImpl) GetToysKidsBabiesProductsByFilter(filters map[string]interface{}, subDept int) ([]*models.ToysKidsBabiesProduct, error) {
	var products []*models.ToysKidsBabiesProduct

	// Start building the query, with Preloads for related tables
	query := GetDbConn().Debug().
		Preload("Company").     // Preload the Company relationship
		Preload("PurchaseInfo") // Preload the PurchaseInfo relationship

	// Apply dynamic product filters
	query = ApplyProductFilters(query, filters, subDept,"toys_kids_babies_products")

	// Execute the query
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

	err := GetDbConn().Create(product).Error
	if err != nil {
		logrus.Errorf("Unable to save product, %s", err)
		return nil, err
	}

	return &product, nil
}

func (s ProductDbServiceImpl) GetCompaniesFromHomeGarden(subDepartment int) ([]*string, error) {
	var companies []*string
	result := GetDbConn().Debug().
		Model(&models.HomeGardenProduct{}).
		Joins("INNER JOIN companies ON companies.id = home_garden_products.company_id").
		Where("home_garden_products.sub_department = ?", subDepartment).
		Distinct("companies.name").
		Pluck("companies.name", &companies)

	if result.Error != nil {
		return nil, result.Error
	}

	logrus.Infof("Companies: %v", companies)
	return companies, nil
}

func (s ProductDbServiceImpl) GetCompaniesFromClothingAccessories(subDepartment int) ([]*string, error) {
	var companies []*string
	result := GetDbConn().
		Model(&models.ClothingAccessoriesProduct{}).
		Joins("INNER JOIN companies ON companies.id = clothing_accessories_products.company_id").
		Where("clothing_accessories_products.sub_department = ?", subDepartment).
		Distinct("companies.name").
		Pluck("companies.name", &companies)

	if result.Error != nil {
		return nil, result.Error
	}

	return companies, nil
}

func (s ProductDbServiceImpl) GetCompaniesFromHealthBathBeauty(subDepartment int) ([]*string, error) {
	var companies []*string
	result := GetDbConn().
		Model(&models.HealthBathBeautyProduct{}).
		Joins("INNER JOIN companies ON companies.id = health_bath_beauty_products.company_id").
		Where("health_bath_beauty_products.sub_department = ?", subDepartment).
		Distinct("companies.name").
		Pluck("companies.name", &companies)

	if result.Error != nil {
		return nil, result.Error
	}

	return companies, nil
}

func (s ProductDbServiceImpl) GetCompaniesFromToysKidsBabies(subDepartment int) ([]*string, error) {
	var companies []*string
	result := GetDbConn().
		Model(&models.ToysKidsBabiesProduct{}).
		Joins("INNER JOIN companies ON companies.id = toys_kids_babies_products.company_id").
		Where("toys_kids_babies_products.sub_department = ?", subDepartment).
		Distinct("companies.name").
		Pluck("companies.name", &companies)

	if result.Error != nil {
		return nil, result.Error
	}

	return companies, nil
}

// HomeGarden - Company Certifications
func (s ProductDbServiceImpl) GetCompanyCertificationsFromHomeGarden(subDepartment int) ([]*string, error) {
	var certifications []*string
	result := GetDbConn().
		Model(&models.HomeGardenProduct{}).
		Joins("INNER JOIN company_certifications ON company_certifications.company_id = home_garden_products.company_id").
		Joins("INNER JOIN certifications ON certifications.id = company_certifications.certification_id").
		Where("home_garden_products.sub_department = ?", subDepartment).
		Distinct("certifications.name").
		Pluck("certifications.name", &certifications)

	if result.Error != nil {
		return nil, result.Error
	}

	return certifications, nil
}

// HealthBathBeauty - Company Certifications
func (s ProductDbServiceImpl) GetCompanyCertificationsFromHealthBathBeauty(subDepartment int) ([]*string, error) {
	var certifications []*string
	result := GetDbConn().
		Model(&models.HealthBathBeautyProduct{}).
		Joins("INNER JOIN company_certifications ON company_certifications.company_id = health_bath_beauty_products.company_id").
		Joins("INNER JOIN certifications ON certifications.id = company_certifications.certification_id").
		Where("health_bath_beauty_products.sub_department = ?", subDepartment).
		Distinct("certifications.name").
		Pluck("certifications.name", &certifications)

	if result.Error != nil {
		return nil, result.Error
	}

	return certifications, nil
}

// ClothingAccessories - Company Certifications
func (s ProductDbServiceImpl) GetCompanyCertificationsFromClothingAccessories(subDepartment int) ([]*string, error) {
	var certifications []*string
	result := GetDbConn().
		Model(&models.ClothingAccessoriesProduct{}).
		Joins("INNER JOIN company_certifications ON company_certifications.company_id = clothing_accessories_products.company_id").
		Joins("INNER JOIN certifications ON certifications.id = company_certifications.certification_id").
		Where("clothing_accessories_products.sub_department = ?", subDepartment).
		Distinct("certifications.name").
		Pluck("certifications.name", &certifications)

	if result.Error != nil {
		return nil, result.Error
	}

	return certifications, nil
}

// ToysKidsBabies - Company Certifications
func (s ProductDbServiceImpl) GetCompanyCertificationsFromToysKidsBabies(subDepartment int) ([]*string, error) {
	var certifications []*string
	result := GetDbConn().
		Model(&models.ToysKidsBabiesProduct{}).
		Joins("INNER JOIN company_certifications ON company_certifications.company_id = toys_kids_babies_products.company_id").
		Joins("INNER JOIN certifications ON certifications.id = company_certifications.certification_id").
		Where("toys_kids_babies_products.sub_department = ?", subDepartment).
		Distinct("certifications.name").
		Pluck("certifications.name", &certifications)

	if result.Error != nil {
		return nil, result.Error
	}

	return certifications, nil
}

// HomeGarden - Product Certifications
func (s ProductDbServiceImpl) GetProductCertificationsFromHomeGarden(subDepartment int) ([]*string, error) {
	var certifications []*string
	result := GetDbConn().
		Model(&models.HomeGardenProduct{}).
		Joins("INNER JOIN product_certifications ON product_certifications.product_id = home_garden_products.id").
		Joins("INNER JOIN certifications ON certifications.id = product_certifications.certification_id").
		Where("home_garden_products.sub_department = ?", subDepartment).
		Distinct("certifications.name").
		Pluck("certifications.name", &certifications)

	if result.Error != nil {
		return nil, result.Error
	}

	return certifications, nil
}

// HealthBathBeauty - Product Certifications
func (s ProductDbServiceImpl) GetProductCertificationsFromHealthBathBeauty(subDepartment int) ([]*string, error) {
	var certifications []*string
	result := GetDbConn().
		Model(&models.HealthBathBeautyProduct{}).
		Joins("INNER JOIN product_certifications ON product_certifications.product_id = health_bath_beauty_products.id").
		Joins("INNER JOIN certifications ON certifications.id = product_certifications.certification_id").
		Where("health_bath_beauty_products.sub_department = ?", subDepartment).
		Distinct("certifications.name").
		Pluck("certifications.name", &certifications)

	if result.Error != nil {
		return nil, result.Error
	}

	return certifications, nil
}

// ClothingAccessories - Product Certifications
func (s ProductDbServiceImpl) GetProductCertificationsFromClothingAccessories(subDepartment int) ([]*string, error) {
	var certifications []*string
	result := GetDbConn().
		Model(&models.ClothingAccessoriesProduct{}).
		Joins("INNER JOIN product_certifications ON product_certifications.product_id = clothing_accessories_products.id").
		Joins("INNER JOIN certifications ON certifications.id = product_certifications.certification_id").
		Where("clothing_accessories_products.sub_department = ?", subDepartment).
		Distinct("certifications.name").
		Pluck("certifications.name", &certifications)

	if result.Error != nil {
		return nil, result.Error
	}

	return certifications, nil
}

// ToysKidsBabies - Product Certifications
func (s ProductDbServiceImpl) GetProductCertificationsFromToysKidsBabies(subDepartment int) ([]*string, error) {
	var certifications []*string
	result := GetDbConn().
		Model(&models.ToysKidsBabiesProduct{}).
		Joins("INNER JOIN product_certifications ON product_certifications.product_id = toys_kids_babies_products.id").
		Joins("INNER JOIN certifications ON certifications.id = product_certifications.certification_id").
		Where("toys_kids_babies_products.sub_department = ?", subDepartment).
		Distinct("certifications.name").
		Pluck("certifications.name", &certifications)

	if result.Error != nil {
		return nil, result.Error
	}

	return certifications, nil
}

// HomeGarden - Price Range
func (s ProductDbServiceImpl) GetPriceRangeFromHomeGarden(subDepartment int) (*model.PriceRange, error) {
	var priceRange model.PriceRange
	result := GetDbConn().
		Model(&models.PurchaseInfo{}).
		Joins("INNER JOIN home_garden_products ON home_garden_products.id = purchase_infos.product_id").
		Where("home_garden_products.sub_department = ?", subDepartment).
		Select("MIN(CAST(purchase_infos.price AS FLOAT)) AS min, MAX(CAST(purchase_infos.price AS FLOAT)) AS max").
		Scan(&priceRange)

	if result.Error != nil {
		return nil, result.Error
	}

	return &priceRange, nil
}

func (s ProductDbServiceImpl) GetPriceRangeFromHealthBathBeauty(subDepartment int) (*model.PriceRange, error) {
	var priceRange model.PriceRange
	result := GetDbConn().
		Model(&models.PurchaseInfo{}).
		Joins("INNER JOIN health_bath_beauty_products ON health_bath_beauty_products.id = purchase_infos.product_id").
		Where("health_bath_beauty_products.sub_department = ?", subDepartment).
		Select("MIN(CAST(purchase_infos.price AS FLOAT)) AS min, MAX(CAST(purchase_infos.price AS FLOAT)) AS max").
		Scan(&priceRange)

	if result.Error != nil {
		return nil, result.Error
	}

	return &priceRange, nil
}

func (s ProductDbServiceImpl) GetPriceRangeFromClothingAccessories(subDepartment int) (*model.PriceRange, error) {
	var priceRange model.PriceRange
	result := GetDbConn().
		Model(&models.PurchaseInfo{}).
		Joins("INNER JOIN clothing_accessories_products ON clothing_accessories_products.id = purchase_infos.product_id").
		Where("clothing_accessories_products.sub_department = ?", subDepartment).
		Select("MIN(CAST(purchase_infos.price AS FLOAT)) AS min, MAX(CAST(purchase_infos.price AS FLOAT)) AS max").
		Scan(&priceRange)

	if result.Error != nil {
		return nil, result.Error
	}

	return &priceRange, nil
}

func (s ProductDbServiceImpl) GetPriceRangeFromToysKidsBabies(subDepartment int) (*model.PriceRange, error) {
	var priceRange model.PriceRange
	result := GetDbConn().
		Model(&models.PurchaseInfo{}).
		Joins("INNER JOIN toys_kids_babies_products ON toys_kids_babies_products.id = purchase_infos.product_id").
		Where("toys_kids_babies_products.sub_department = ?", subDepartment).
		Select("MIN(CAST(purchase_infos.price AS FLOAT)) AS min, MAX(CAST(purchase_infos.price AS FLOAT)) AS max").
		Scan(&priceRange)

	if result.Error != nil {
		return nil, result.Error
	}

	return &priceRange, nil
}
