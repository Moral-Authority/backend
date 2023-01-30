package handlers

import (
	"errors"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
)

type ProductCategorizationService struct{}

func (s ProductService) AddNewSection(request model.AddSection, dbService database.ProductDbService) (*model.Section, error) {
	department := models.Section{
		Title:      request.Title,
		Categories: []models.Category{},
	}
	addedDept := dbService.AddSection(department)
	if addedDept == nil {
		return nil, errors.New("unable to save department to db")
	}
	return toSectionResponse(*addedDept), nil
}

func (s ProductService) AddNewSubSection(request model.AddSubSection, dbService database.ProductDbService) (*model.SubSection, error) {
	department := models.Department{
		Title:      request.Title,
		Categories: []models.Category{},
	}
	addedDept := dbService.AddSubSection(department)
	if addedDept == nil {
		return nil, errors.New("unable to save department to db")
	}
	return toSubSectionResponse(*addedDept), nil
}

func (s ProductService) AddNewDepartment(request model.AddDepartment, dbService database.ProductDbService) (*model.Department, error) {
	department := models.Department{
		Title:      request.Title,
		Categories: []models.Category{},
	}
	addedDept := dbService.AddDepartment(department)
	if addedDept == nil {
		return nil, errors.New("unable to save department to db")
	}
	return toDepartmentResponse(*addedDept), nil
}

func (s ProductService) AddNewCategory(request model.AddCategory, dbService database.ProductDbService) (*model.Category, error) {
	deptId, err := database.StringToUint(request.DepartmentID)
	if err != nil {
		return nil, err
	}
	category := models.Category{
		Title:           request.Title,
		DepartmentRefer: deptId,
		Types:           []models.Type{},
	}
	addedCategory := dbService.AddCategory(category)
	if addedCategory == nil {
		return nil, errors.New("unable to save category to db")
	}
	return toCategoryResponse(*addedCategory), nil
}

func (s ProductService) AddNewSubCategory(request model.AddSubCategory, dbService database.ProductDbService) (*model.SubCategory, error) {
	categoryID, err := database.StringToUint(request.CategoryID)
	if err != nil {
		return nil, err
	}
	subcategory := models.SubCategory{
		Title:         request.Title,
		CategoryRefer: categoryID,
	}
	addedCategory := dbService.AddSubCategory(subcategory)
	if addedCategory == nil {
		return nil, errors.New("unable to save category to db")
	}
	return toSubCategoryResponse(*addedCategory), nil
}

func (s ProductService) AddNewType(request model.AddTypeRequest, dbService database.ProductDbService) (*model.Type, error) {
	catId, err := database.StringToUint(request.SubCategoryID)
	if err != nil {
		return nil, err
	}
	typeModel := models.Type{
		Title:         request.Title,
		CategoryRefer: catId,
		Styles:        []models.Style{},
	}
	addedType := dbService.AddType(typeModel)
	if addedType == nil {
		return nil, errors.New("unable to save type to db")
	}
	return toTypeResponse(*addedType), nil
}

func (s ProductService) AddNewStyle(request model.AddStyleRequest, dbService database.ProductDbService) (*model.Style, error) {
	typeId, err := database.StringToUint(request.TypeID)
	if err != nil {
		return nil, err
	}
	style := models.Style{
		Title:     request.Title,
		TypeRefer: typeId,
		Products:  []models.Product{},
	}
	addedStyle := dbService.AddStyle(style)
	if addedStyle == nil {
		return nil, errors.New("unable to save style to db")
	}
	return toStyleResponse(*addedStyle), nil
}

func (s ProductService) AddNewFilter(request model.AddStyleRequest, dbService database.ProductDbService) (*model.Style, error) {
	typeId, err := database.StringToUint(request.TypeID)
	if err != nil {
		return nil, err
	}
	style := models.Style{
		Title:     request.Title,
		TypeRefer: typeId,
		Products:  []models.Product{},
	}
	addedStyle := dbService.AddStyle(style)
	if addedStyle == nil {
		return nil, errors.New("unable to save style to db")
	}
	return toStyleResponse(*addedStyle), nil
}

func (s ProductService) GetDepartments(dbService database.ProductDbService) ([]*model.Department, error) {
	departments := dbService.GetDepartments()
	if departments == nil {
		return nil, errors.New("unable to get departments from db")
	}
	return toDepartmentsResponse(departments), nil
}
