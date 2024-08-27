package database

import (
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
)

type ProductDbServiceImpl struct{}

func (s ProductDbServiceImpl) GetProductByID(productId string) (*models.Product, error) {
	var product models.Product

	result := GetDbConn().First(&product, "id = ?", productId)
	if result.Error != nil {
		logrus.Errorf("Unable to get product, %s", result.Error)
		return nil, result.Error
	}

	return &product, nil
}

func (s ProductDbServiceImpl) GetAllProducts() ([]*models.Product, error) {
	var products []*models.Product
	result := GetDbConn().Find(&products)
	if result.Error != nil {
		logrus.Errorf("Unable to get all products, %s", result.Error)
		return nil, result.Error
	}
	return products, nil
}

func (s ProductDbServiceImpl) GetProductsByFilter(filters map[string]interface{}) ([]models.Product, error) {
	var products []models.Product
	db := GetDbConn()

	query := ApplyFilters(db, filters)

	if err := query.Find(&products).Error; err != nil {
		logrus.Errorf("Unable to get certifications by filter, %s", err)
		return nil, err
	}

	return products, nil
}

func (s ProductDbServiceImpl) AddProduct(product models.Product) (*models.Product, error) {
	result := GetDbConn().Create(&product)
	if result.Error != nil {
		logrus.Errorf("Unable to save product, %s", result.Error)
		return nil, result.Error
	}
	var addedProduct models.Product
	result = GetDbConn().First(&addedProduct, "id = ?", product.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get product, %s", result.Error)
		return nil, result.Error
	}
	return &addedProduct, nil
}

func (s ProductDbServiceImpl) AddProductCertification(pc models.ProductCertification) (*models.ProductCertification, error) {
	result := GetDbConn().Create(&pc)
	if result.Error != nil {
		logrus.Errorf("Unable to save product, %s", result.Error)
		return nil, result.Error
	}
	var added models.ProductCertification
	result = GetDbConn().First(&added, "id = ?", pc.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get product, %s", result.Error)
		return nil, result.Error
	}
	return &added, nil
}

func (s ProductDbServiceImpl) UpdateProduct(product model.UpdateProductRequest) (*models.Product, error) {

	return nil, nil
}

func (s ProductDbServiceImpl) DeleteProduct(productId string) error {
	var product models.Product
	result := GetDbConn().First(&product, "id = ?", productId)
	if result.Error != nil {
		logrus.Errorf("Unable to find product, %s", result.Error)
		return result.Error
	}
	result = GetDbConn().Delete(&product)
	if result.Error != nil {
		logrus.Errorf("Unable to delete product, %s", result.Error)
		return result.Error
	}
	return nil
}

func (s ProductDbServiceImpl) AddCategory(category models.Category) (*models.Category, error) {
	result := GetDbConn().Create(&category)
	if result.Error != nil {
		logrus.Errorf("Unable to save category, %s", result.Error)
		return nil, result.Error
	}
	var addedCategory models.Category
	result = GetDbConn().First(&addedCategory, "id = ?", category.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get category, %s", result.Error)
		return nil, result.Error
	}
	return &addedCategory, nil
}

func (s ProductDbServiceImpl) GetAllCategories() ([]*models.Category, error) {
	var categories []*models.Category
	result := GetDbConn().Find(&categories)
	if result.Error != nil {
		logrus.Errorf("Unable to get all categories, %s", result.Error)
		return nil, result.Error
	}
	return categories, nil
}

//var response []*models.Category
//for _, dept := range categories {
//	var categories []models.Category
//	result := GetDbConn().Model(&dept).Association("Categories").Find(&categories)
//	if result != nil {
//		logrus.Errorf("Unable to get categories, %s", result)
//		return nil
//	}
//for _, cat := range categories {
//	var types []models.Category
//	result := GetDbConn().Model(&cat).Association("Types").Find(&types)
//	if result != nil {
//		logrus.Errorf("Unable to get types, %s", result)
//		return nil
//	}
//	for _, type_ := range types {
//		var styles []models.Style
//		result := GetDbConn().Model(&type_).Association("Styles").Find(&styles)
//		if result != nil {
//			logrus.Errorf("Unable to get styles, %s", result)
//			return nil
//		}
//		for _, style := range styles {
//			var products []models.Product
//			result := GetDbConn().Model(&style).Association("Products").Find(&products)
//			if result != nil {
//				logrus.Errorf("Unable to get products, %s", result)
//				return nil
//			}
//			style.Products = products
//		}
//		type_.Styles = styles
//	}
//	cat.Types = types
//}
//dept.Categories = categories
//	response = append(response, dept)
//}
