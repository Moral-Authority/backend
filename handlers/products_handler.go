package handlers

import (
	"fmt"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
)

type ProductService struct{}

// func (s ProductService) AddNewProductHandler(request model.AddProductRequest, productDbService database.ProductDbService, imageDbService database.ImageDbService, certificationService database.CertificationDbService) (*model.Product, error) {

// 	var product interface{}
// 	productBase := models.ProductBase{
// 		Url:         request.PurchaseInfo.Link,
// 		Description: request.Description,
// 		Title:       request.Title,
// 	}

// 	pd, isDepartment := IsStringValidProductDepartment(request.Department)
// 	if !isDepartment {
// 		return nil, fmt.Errorf("invalid department type: %s", request.Department)
// 	}

// 	switch pd {
// 	case HomeGardenProductDepartment:
// 		product = &models.HomeGardenProduct{
// 			ProductBase: productBase,
// 		}
// 	case BathBeautyProductDepartment:
// 		product = &models.HealthBathBeautyProduct{
// 			ProductBase: productBase,
// 		}
// 	case ClothingAccessoriesProductDepartment:
// 		product = &models.ClothingAccessoriesProduct{
// 			ProductBase: productBase,
// 		}
// 	case ToysKidsBabiesProductDepartment:
// 		product = &models.ToysKidsBabiesProduct{
// 			ProductBase: productBase,
// 		}
// 	default:
// 		return nil, fmt.Errorf("unknown department type: %d", request.Department)
// 	}

// 	// Call the helper function to add the product by department
// 	return AddProductByDepartment(pd, product, request, productDbService, imageDbService, certificationService)
// }

func (s ProductService) UpdateProductHandler(request model.UpdateProductRequest, productDbService database.ProductDbService, imageDbService database.ImageDbService, certificationService database.CertificationDbService) (*model.Product, error) {
	product, err := productDbService.UpdateProduct(request)
	if err != nil {
		return nil, err
	}

	productDept, isDepartment := IsStringValidProductDepartment(*request.Department)
	if !isDepartment {
		return nil, fmt.Errorf("invalid department type: %s", request.Department)
	}

	return toProductResponse(product, productDept), nil
}

func (s ProductService) GetProductByIDHandler(productId string, department int, productDbService database.ProductDbService) (*model.Product, error) {

	_, isDepartment := IsValidProductDepartment(department)
	if !isDepartment {
		return nil, fmt.Errorf("invalid department type: %s", department)
	}

	productDept := ProductDepartment(department)
	var product interface{}
	var err error

	switch productDept {
	case HomeGardenProductDepartment:
		product, err = productDbService.GetHomeGardenProductByID(productId)
	case BathBeautyProductDepartment:
		product, err = productDbService.GetBathBeautyProductByID(productId)
	case ClothingAccessoriesProductDepartment:
		product, err = productDbService.GetClothingAccessoriesProductByID(productId)
	case ToysKidsBabiesProductDepartment:
		product, err = productDbService.GetToysKidsBabiesProductByID(productId)
	default:
		return nil, fmt.Errorf("unknown department type: %d", department)
	}

	if err != nil {
		return nil, err
	}

	return toProductResponse(product, productDept), nil
}

func (s ProductService) GetAllProductsHandler(productDbService database.ProductDbService, department int) ([]*model.Product, error) {

	productDept, isDepartment := IsValidProductDepartment(department)
	if !isDepartment {
		return nil, fmt.Errorf("invalid department type: %s", department)
	}

	switch productDept {
	case HomeGardenProductDepartment:
		products, err := productDbService.GetAllHomeGardenProducts()
		if err != nil {
			return nil, err
		}
		return toHomeGardenProductsResponse(products, productDept), nil
	case BathBeautyProductDepartment:
		products, err := productDbService.GetAllBathBeautyProducts()
		if err != nil {
			return nil, err
		}
		return toBathBeautyProductsResponse(products, productDept), nil
	case ClothingAccessoriesProductDepartment:
		products, err := productDbService.GetAllClothingAccessoriesProducts()
		if err != nil {
			return nil, err
		}
		return toClothingAccessoriesProductsResponse(products, productDept), nil
	case ToysKidsBabiesProductDepartment:
		products, err := productDbService.GetAllToysKidsBabiesProducts()
		if err != nil {
			return nil, err
		}
		return toToysKidsBabiesProductsResponse(products, productDept), nil
	default:
		return nil, fmt.Errorf("unknown department type: %d", department)
	}
}

func (s ProductService) GetProductsByFilterHandler(productDbService database.ProductDbService, filters map[string]interface{}, department int) ([]*model.Product, error) {

	productDept, isDepartment := IsValidProductDepartment(department)
	if !isDepartment {
		return nil, fmt.Errorf("invalid department type: %s", department)
	}

	switch productDept {
	case HomeGardenProductDepartment:
		products, err := productDbService.GetHomeGardenProductsByFilter(filters)
		if err != nil {
			return nil, err
		}
		return toHomeGardenProductsResponse(products, productDept), nil

	case BathBeautyProductDepartment:
		products, err := productDbService.GetBathBeautyProductsByFilter(filters)
		if err != nil {
			return nil, err
		}
		return toBathBeautyProductsResponse(products, productDept), nil

	case ClothingAccessoriesProductDepartment:
		products, err := productDbService.GetClothingAccessoriesProductsByFilter(filters)
		if err != nil {
			return nil, err
		}
		return toClothingAccessoriesProductsResponse(products, productDept), nil

	case ToysKidsBabiesProductDepartment:
		products, err := productDbService.GetToysKidsBabiesProductsByFilter(filters)
		if err != nil {
			return nil, err
		}
		return toToysKidsBabiesProductsResponse(products, productDept), nil
	default:
		return nil, fmt.Errorf("unknown department type: %d", department)
	}
}
