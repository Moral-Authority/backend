package handlers

import (
	"fmt"
	"strconv"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/volatiletech/null/v8"
)

type ProductService struct{}

// func (s ProductService) AddNewProductHandler(request model.AddProductRequest, productDbService database.ProductDbService, imageDbService database.ImageDbService, certificationService database.CertificationDbService, companyService database.CompanyDbServiceImpl) (*model.Product, error) {

// 	var product interface{}
// 	productBase := models.ProductBase{
// 		Url:           request.PurchaseInfo.Link,
// 		Description:   request.Description,
// 		Title:         request.Title,
// 		ProductImage:  *request.ImageLinks[0],
// 	}

// 	prodDept, isValid := IsStringValidProductDepartment(request.Department)
// 	if !isValid {
// 		return nil, fmt.Errorf("invalid department type: %s", request.Department)
// 	}

// 	subDept, isValid := IsStringValidProductSubDepartmentFORSEED(prodDept, request.SubDepartment)
// 	if !isValid {
// 		return nil, fmt.Errorf("invalid department type: %s", request.Department)
// 	}

// 	companyID, err := companyService.FindCompanyIDByName(request.Company)
// 	if err != nil {
// 		return nil, err
// 	}

// 	productCertIDs := []uint{}
// 	for _, cert := range request.ProductCertifications {
// 		id, err := certificationService.GetCertificationIdByName(cert)
// 		if err != nil {
// 			return nil, err
// 		}
// 		productCertIDs = append(productCertIDs, id)
// 	}

// 	companyCertsIDs := []uint{}
// 	for _, cert := range request.CompanyCertifications {
// 		id, err := certificationService.GetCertificationIdByName(cert)
// 		if err != nil {
// 			return nil, err
// 		}
// 		companyCertsIDs = append(companyCertsIDs, id)
// 	}

// 	productBase.SubDepartment = subDept
// 	productBase.CompanyID = companyID

// 	switch prodDept {
// 	case HomeGardenProductDepartment:
// 		product = &models.HomeGardenProduct{
// 			ProductBase: productBase,
// 		}
// 	case HealthBathBeautyProductDepartment:
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

// 	for _, certID := range productCertIDs {
// 		prodCert := models.ProductCertification{
// 			CertificationID: certID,
// 			ProductID: 	 product.(models.Product).GetID(),
// 		}
// 		result := db.Create(&prodCert)
// 		if result.Error != nil {
// 			fmt.Printf("error inserting ProductCertification for product: %v", result.Error)
// 		}
// 	}

// 	for _, certID := range companyCertsIDs {
// 		compCert := models.CompanyCertification{
// 			CertificationID: certID,
// 			CompanyID: 	 companyID,
// 		}
// 		result := db.Create(&compCert)

// 		if result.Error != nil {
// 			fmt.Printf("error inserting CompanyCertification for product: %v", result.Error)
// 		}
// 	}

// 	price, err := strconv.ParseFloat(row[4], 64)
// 	if err != nil {
// 		fmt.Printf("invalid price format: %v", err)
// 	}

// 	purchaseInfo := models.PurchaseInfo{
// 		ProductID:         productID,
// 		ProductDepartment: prodDeptInt,
// 		Price:             price,
// 		Url:               row[5],
// 	}

// 	result := db.Create(&purchaseInfo)
// 	if result.Error != nil {
// 		fmt.Printf("error inserting PurchaseInfo for ToysKidsBabies product: %v", result.Error)
// 	}

// 	// Index the product in Algolia
// 	algoliaData := map[string]interface{}{
// 		"objectID":       productID,
// 		"title":          row[3],
// 		"sub_department": row[1],
// 		"url":            row[5],
// 		"company_name":   companyName,
// 		"product_image":  row[6],
// 		"description":    row[11],
// 		"price":          row[4],
// 		"department":     row[0],
// 	}

// 	_, err = index.SaveObject(algoliaData)
// 	if err != nil {
// 		fmt.Printf("Failed to index product in Algolia: %v", err)
// 	} else {
// 		fmt.Printf("Indexed product in Algolia: %s\n", row[4])
// 	}

// 	// Call the helper function to add the product by department
// 	return toProductResponse(product, productDept), nil
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


func (s ProductService) GetRecentlyAddedProductsHandler(productDbService database.ProductDbService) ([]*model.Product, error) {
	products, err := productDbService.GetRecentlyAddedProducts()
	if err != nil {
		return nil, err
	}

	return toHomeGardenProductsResponse(products, HomeGardenProductDepartment), nil
}