package handlers

import (
	"strconv"

	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
)

func toProductResponse(product interface{}, department ProductDepartment) *model.Product {
	var baseProduct *models.ProductBase

	// Use type assertion for each product type
	switch p := product.(type) {
	case *models.HomeGardenProduct:
		baseProduct = &p.ProductBase
	case *models.HealthBathBeautyProduct:
		baseProduct = &p.ProductBase
	case *models.ClothingAccessoriesProduct:
		baseProduct = &p.ProductBase
	case *models.ToysKidsBabiesProduct:
		baseProduct = &p.ProductBase
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
		// ProductCertifications: toCertificationsResponse(baseProduct.ProductCertifications),
		Department: department.ToString(),
	}
}


func toGenericProductsResponse[T any](products []*T, department ProductDepartment) []*model.Product {
	var response []*model.Product

	for _, e := range products {
		product := toProductResponse(e, department)
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
