package database

import (
	"github.com/howstrongiam/backend/models"
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

func (s ProductDbServiceImpl) AddDepartment(department models.Department) *models.Department {
	result := GetDbConn().Create(&department)
	if result.Error != nil {
		logrus.Errorf("Unable to save department, %s", result.Error)
		return nil
	}
	var addedDepartment models.Department
	result = GetDbConn().First(&addedDepartment, "id = ?", department.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get department, %s", result.Error)
		return nil
	}
	return &addedDepartment
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

func (s ProductDbServiceImpl) AddType(typeToAdd models.Type) *models.Type {
	result := GetDbConn().Create(&typeToAdd)
	if result.Error != nil {
		logrus.Errorf("Unable to save type, %s", result.Error)
		return nil
	}
	var addedType models.Type
	result = GetDbConn().First(&addedType, "id = ?", typeToAdd.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get type, %s", result.Error)
		return nil
	}
	return &addedType
}

func (s ProductDbServiceImpl) AddStyle(style models.Style) *models.Style {
	result := GetDbConn().Create(&style)
	if result.Error != nil {
		logrus.Errorf("Unable to save style, %s", result.Error)
		return nil
	}
	var addedStyle models.Style
	result = GetDbConn().First(&addedStyle, "id = ?", style.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get style, %s", result.Error)
		return nil
	}
	return &addedStyle
}

func (s ProductDbServiceImpl) AddProduct(product models.Product) *models.Product {
	result := GetDbConn().Create(&product)
	if result.Error != nil {
		logrus.Errorf("Unable to add product, %s", result.Error)
		return nil
	}
	var addedProduct models.Product
	result = GetDbConn().First(&addedProduct, "id = ?", product.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get product, %s", result.Error)
		return nil
	}
	return &addedProduct
}

func (s ProductDbServiceImpl) GetDepartments() []*models.Department {
	var departments []*models.Department
	result := GetDbConn().Find(&departments)
	if result.Error != nil {
		logrus.Errorf("Unable to get all departments, %s", result.Error)
		return nil
	}
	var response []*models.Department
	for _, dept := range departments {
		var categories []models.Category
		result := GetDbConn().Model(&dept).Association("Categories").Find(&categories)
		if result != nil {
			logrus.Errorf("Unable to get categories, %s", result)
			return nil
		}
		for _, cat := range categories {
			var types []models.Type
			result := GetDbConn().Model(&cat).Association("Types").Find(&types)
			if result != nil {
				logrus.Errorf("Unable to get types, %s", result)
				return nil
			}
			for _, type_ := range types {
				var styles []models.Style
				result := GetDbConn().Model(&type_).Association("Styles").Find(&styles)
				if result != nil {
					logrus.Errorf("Unable to get styles, %s", result)
					return nil
				}
				for _, style := range styles {
					var products []models.Product
					result := GetDbConn().Model(&style).Association("Products").Find(&products)
					if result != nil {
						logrus.Errorf("Unable to get products, %s", result)
						return nil
					}
					style.Products = products
				}
				type_.Styles = styles
			}
			cat.Types = types
		}
		dept.Categories = categories
		response = append(response, dept)
	}
	return response
}
