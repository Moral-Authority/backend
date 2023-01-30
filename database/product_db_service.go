package database

import "github.com/howstrongiam/backend/models"

type ProductDbService interface {
	GetProduct(productId string) *models.Product

	//AddSection(department models.Section) *models.Section
	//AddSubSection(department models.SubSection) *models.SubSection
	AddDepartment(department models.Department) *models.Department
	AddCategory(category models.Category) *models.Category
	//AddSubCategory(category models.SubCategory) *models.SubCategory
	AddType(typeToAdd models.Type) *models.Type
	AddStyle(style models.Style) *models.Style
	AddProduct(product models.Product) *models.Product
	GetDepartments() []*models.Department
}
