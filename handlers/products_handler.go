package handlers

import (
	"errors"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
)

type ProductService struct{}

func (s ProductService) AddNewDepartment(request model.AddDepartmentRequest, dbService database.ProductDbService) (*model.Department, error) {
	department := models.Department{
		Title:      request.Title,
		Categories: []models.Category{},
	}
	addedDept := dbService.AddDepartment(department)
	if addedDept == nil {
		return nil, errors.New("unable to save department to db")
	}
	return nil, nil
}

func (s ProductService) AddNewCategory(request model.AddCategoryRequest, dbService database.ProductDbService) (*model.Category, error) {
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
	return nil, nil
}

func (s ProductService) AddNewType(request model.AddTypeRequest, dbService database.ProductDbService) (*model.Type, error) {
	catId, err := database.StringToUint(request.CategoryID)
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
	return nil, nil
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
	return nil, nil
}

func (s ProductService) AddNewProduct(
	request model.AddProductRequest,
	productDbService database.ProductDbService,
	imageDbService database.ImageDbService,
	certDbService database.CertificatesDbService,
) (*model.Product, error) {
	styleId, err := database.StringToUint(request.StyleID)
	if err != nil {
		return nil, err
	}
	userId, err := database.StringToUint(request.UserID)
	if err != nil {
		return nil, err
	}
	image := models.Image{
		ImageLocation: request.ImageLocation,
	}
	addedImage := imageDbService.AddImage(image)
	if addedImage == nil {
		return nil, errors.New("unable to save image to db")
	}
	cert := models.Certification{
		CertifyingCompany: request.Certification.CertifyingCompany,
		CertName:          request.Certification.CertName,
	}
	addedCert := certDbService.AddNewCertificate(cert)
	if addedCert == nil {
		return nil, errors.New("unable to save certificate to db")
	}
	product := models.Product{
		Url:             request.URL,
		Description:     request.Description,
		Title:           request.Title,
		UserId:          userId,
		ImageId:         addedImage.ID,
		CertificationId: addedCert.ID,
		StyleRefer:      styleId,
	}
	addedProduct := productDbService.AddProduct(product)
	if addedProduct == nil {
		return nil, errors.New("unable to save product to db")
	}
	return nil, nil
}

func (s ProductService) GetDepartments(dbService database.ProductDbService) ([]*model.Department, error) {
	departments := dbService.GetDepartments()
	if departments == nil {
		return nil, errors.New("unable to get departments")
	}
	return toDepartmentsResponse(departments), nil
}
