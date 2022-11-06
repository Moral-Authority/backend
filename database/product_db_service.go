package database

import "github.com/howstrongiam/backend/models"

type ProductDbService interface {
	GetProduct(productId string) *models.Product
	AddDepartment(department models.Department) *models.Department
	AddCategory(category models.Category) *models.Category
	AddType(typeToAdd models.Type) *models.Type
	AddStyle(style models.Style) *models.Style
	AddProduct(product models.Product) *models.Product
	GetDepartments() []*models.Department
}
