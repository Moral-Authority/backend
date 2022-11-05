package handlers

import "github.com/howstrongiam/backend/graph/model"

type ProductService struct{}

func (s ProductService) AddNewDepartment(request model.AddDepartmentRequest) (*model.Department, error) {
	panic("NOT IMPL")
}

func (s ProductService) AddNewType(request model.AddTypeRequest) (*model.Type, error) {
	panic("NOT IMPL")
}

func (s ProductService) AddNewStyle(request model.AddStyleRequest) (*model.Style, error) {
	panic("NOT IMPL")
}

func (s ProductService) AddNewProduct(request model.AddProductRequest) (*model.Product, error) {
	panic("NOT IMPL")
}

func (s ProductService) GetDepartments() ([]*model.Department, error) {
	panic("NOT IMPL")
}
