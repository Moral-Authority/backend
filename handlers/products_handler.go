package handlers

import (
	"fmt"
	"strconv"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
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
		return nil, fmt.Errorf("invalid department type: %s", *request.Department)
	}

	return toProductResponse(product, productDept), nil
}

func (s ProductService) GetProductByIDHandler(productId string, department int, productDbService database.ProductDbService) (*model.Product, error) {

	productDept, isDepartment := IsValidProductDepartment(department)
	if !isDepartment {
		return nil, fmt.Errorf("invalid department type: %s", strconv.Itoa(department))
	}

	prodID, err := database.StringToUint(productId)
	if err != nil {
		return nil, err
	}

	var product interface{}

	switch productDept {
	case HomeGardenProductDepartment:
		product, err = productDbService.GetHomeGardenProductByID(prodID)
	case HealthBathBeautyProductDepartment:
		product, err = productDbService.GetHealthBathBeautyProductByID(prodID)
	case ClothingAccessoriesProductDepartment:
		product, err = productDbService.GetClothingAccessoriesProductByID(prodID)
	case ToysKidsBabiesProductDepartment:
		product, err = productDbService.GetToysKidsBabiesProductByID(prodID)
	default:
		return nil, fmt.Errorf("unknown department type: %d", department)
	}

	if err != nil {
		return nil, err
	}

	return toProductResponse(product, productDept), nil
}

func (s ProductService) GetAllProductsHandler(productDbService database.ProductDbService, productDept ProductDepartment, subDepartment string) ([]*model.Product, error) {

	logrus.Info("Getting all products %s %s", productDept.ToString(), subDepartment)
	switch productDept {
	case HomeGardenProductDepartment:
		subDept, isSubDept := IsStringValidHomeGardenSubDep(subDepartment)
		if !isSubDept {
			return nil, fmt.Errorf("invalid sub-department type: %s", subDepartment)
		}

		products, err := productDbService.GetAllHomeGardenProducts(null.IntFrom(subDept.ToInt()))
		if err != nil {
			return nil, err
		}
		return toHomeGardenProductsResponse(products, productDept), nil

	case HealthBathBeautyProductDepartment:
		subDept, isSubDept := IsStringValidHealthBathBeautySubDep(subDepartment)
		if !isSubDept {
			return nil, fmt.Errorf("invalid sub-department type: %s", subDepartment)
		}
		
		products, err := productDbService.GetAllBathBeautyProducts(null.IntFrom(subDept.ToInt()))
		if err != nil {
			return nil, err
		}
		return toBathBeautyProductsResponse(products, productDept), nil

	case ClothingAccessoriesProductDepartment:
		subDept, isSubDept := IsStringValidClothingAccessoriesSubDep(subDepartment)
		if !isSubDept {
			return nil, fmt.Errorf("invalid sub-department type: %s", subDepartment)
		}
		
		products, err := productDbService.GetAllClothingAccessoriesProducts(null.IntFrom(subDept.ToInt()))
		if err != nil {
			return nil, err
		}
		return toClothingAccessoriesProductsResponse(products, productDept), nil

	case ToysKidsBabiesProductDepartment:
		subDept, isSubDept := IsStringValidToysKidsBabiesSubDep(subDepartment)
		if !isSubDept {
			return nil, fmt.Errorf("invalid sub-department type: %s", subDepartment)
		}
		
		products, err := productDbService.GetAllToysKidsBabiesProducts(null.IntFrom(subDept.ToInt()))
		if err != nil {
			return nil, err
		}
		return toToysKidsBabiesProductsResponse(products, productDept), nil
	default:
		return nil, fmt.Errorf("unknown department type: %d", productDept)
	}
}