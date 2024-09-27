package handlers

import (
	"strconv"

	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
)

func toProductResponse(product interface{}, department ProductDepartment, productCerts []*models.Certification, companyCerts []*models.Certification) *model.Product {
	var baseProduct *models.ProductBase

	// Use type assertion for each product type
	var subDept string
	switch p := product.(type) {
	case *models.HomeGardenProduct:
		baseProduct = &p.ProductBase
		subDept = HomeGardenSubDep(baseProduct.SubDepartment).ToString()
	case *models.HealthBathBeautyProduct:
		baseProduct = &p.ProductBase
		subDept = HealthBathBeautySubDep(baseProduct.SubDepartment).ToString()
	case *models.ClothingAccessoriesProduct:
		baseProduct = &p.ProductBase
		subDept = ClothingAccessoriesSubDep(baseProduct.SubDepartment).ToString()
	case *models.ToysKidsBabiesProduct:
		baseProduct = &p.ProductBase
		subDept = ToysKidsBabiesSubDep(baseProduct.SubDepartment).ToString()
	default:
		logrus.Errorf("Invalid product type: expected *models.ProductBase, got %T", product)
		return nil
	}

	return &model.Product{
		ID:           strconv.Itoa(int(baseProduct.ID)),
		Title:        baseProduct.Title,
		Description:  baseProduct.Description,
		ImageLinks:   []string{baseProduct.ProductImage},
		Company:      toCompanyResponse(&baseProduct.Company),
		PurchaseInfo: toPurchaseInfoResponse(baseProduct.PurchaseInfo),
		SubDepartment: subDept,
		ProductCertifications: toCertificationsResponse(productCerts),
		CompanyCertifications: toCertificationsResponse(companyCerts),
		Department: department.ToString(),
	}
}

func toGenericProductsResponse[T any](products []*T, department ProductDepartment) []*model.Product {
	var response []*model.Product

	for _, e := range products {
		product := toProductResponse(e, department, nil, nil)
		response = append(response, product)
	}
	return response
}

func toHomeGardenProductsResponse(products []*models.HomeGardenProduct, department ProductDepartment) []*model.Product {
	return toGenericProductsResponse(products, department)
}

func toBathBeautyProductsResponse(products []*models.HealthBathBeautyProduct, department ProductDepartment) []*model.Product {
	return toGenericProductsResponse(products, department)
}

func toClothingAccessoriesProductsResponse(products []*models.ClothingAccessoriesProduct, department ProductDepartment) []*model.Product {
	return toGenericProductsResponse(products, department)
}

func toToysKidsBabiesProductsResponse(products []*models.ToysKidsBabiesProduct, department ProductDepartment) []*model.Product {
	return toGenericProductsResponse(products, department)
}
