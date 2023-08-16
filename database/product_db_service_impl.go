package database

import (
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
)

type ProductDbServiceImpl struct{}

func (s ProductDbServiceImpl) GetProduct(productId string) *models.Product {
	var product models.Product
	result := GetDbConn().First(&product, "id = ?", productId)
	if result.Error != nil {
		logrus.Errorf("Unable to get product, %s", result.Error)
		return nil
	}
	return &product
}

func (s ProductDbServiceImpl) AddCategory(category models.Category) *models.Category {
	result := GetDbConn().Create(&category)
	if result.Error != nil {
		logrus.Errorf("Unable to save category, %s", result.Error)
		return nil
	}
	var addedCategory models.Category
	result = GetDbConn().First(&addedCategory, "id = ?", category.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get category, %s", result.Error)
		return nil
	}
	return &addedCategory
}

func (s ProductDbServiceImpl) GetAllCategories() []*models.Category {
	var categories []*models.Category
	result := GetDbConn().Find(&categories)
	if result.Error != nil {
		logrus.Errorf("Unable to get all categories, %s", result.Error)
		return nil
	}
	return categories
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
