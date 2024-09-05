package handlers

import (
	"context"
	"fmt"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
)


func (s ProductService) GetSubDepartmentFiltersHandler(ctx context.Context, department string, subDepartment string, productDbService database.ProductDbService) (*model.Filters, error) {
	var companies []*string
	var companyCertifications []*string
	var productCertifications []*string
	var priceRange *model.PriceRange
	var err error

	// Default values
	defaultCompanies := []*string{ConvertString("Adidas"), ConvertString("Nike"), ConvertString("Puma")}
	defaultCompanyCertifications := []*string{
		ConvertString("BCorp"),
		ConvertString("Women Owned"),
		ConvertString("Black Owned"),
		ConvertString("Indigenous Owned"),
		ConvertString("Asian Pacific Islander"),
		ConvertString("LGBTQ+ Owned"),
		ConvertString("Political Donations?"),
	}
	defaultProductCertifications := []*string{
		ConvertString("Fair Trade"),
		ConvertString("Leaping Bunny"),
		ConvertString("Plastic - Free"),
		ConvertString("Vegan"),
		ConvertString("Organic"),
		ConvertString("Made in America"),
	}
	defaultPrice := float64(0.0)
	defaultPriceRange := &model.PriceRange{
		Min: &defaultPrice, 
		Max: &defaultPrice, 
	}

	// Validate department and subDepartment
	prodDeptType, isDept := IsStringValidProductDepartment(department)
	if !isDept {
		fmt.Printf("Invalid product department %s\n", department)
	}

	subDept, isSubdept := IsStringValidProductSubDepartmentFORSEED(prodDeptType, subDepartment)
	if !isSubdept {
		fmt.Printf("Invalid subdepartment %s\n", subDepartment)
	}

	// Query the database based on the department
	switch prodDeptType {
	case HomeGardenProductDepartment:
		companies, err = productDbService.GetCompaniesFromHomeGarden(subDept)
		if err != nil {
			return nil, err
		}
		companyCertifications, err = productDbService.GetCompanyCertificationsFromHomeGarden(subDept)
		if err != nil {
			return nil, err
		}
		productCertifications, err = productDbService.GetProductCertificationsFromHomeGarden(subDept)
		if err != nil {
			return nil, err
		}
		priceRange, err = productDbService.GetPriceRangeFromHomeGarden(subDept)
		if err != nil {
			return nil, err
		}
	case HealthBathBeautyProductDepartment:
		companies, err = productDbService.GetCompaniesFromHealthBathBeauty(subDept)
		if err != nil {
			return nil, err
		}
		companyCertifications, err = productDbService.GetCompanyCertificationsFromHealthBathBeauty(subDept)
		if err != nil {
			return nil, err
		}
		productCertifications, err = productDbService.GetProductCertificationsFromHealthBathBeauty(subDept)
		if err != nil {
			return nil, err
		}
		priceRange, err = productDbService.GetPriceRangeFromHealthBathBeauty(subDept)
		if err != nil {
			return nil, err
		}
	case ClothingAccessoriesProductDepartment:
		companies, err = productDbService.GetCompaniesFromClothingAccessories(subDept)
		if err != nil {
			return nil, err
		}
		companyCertifications, err = productDbService.GetCompanyCertificationsFromClothingAccessories(subDept)
		if err != nil {
			return nil, err
		}
		productCertifications, err = productDbService.GetProductCertificationsFromClothingAccessories(subDept)
		if err != nil {
			return nil, err
		}
		priceRange, err = productDbService.GetPriceRangeFromClothingAccessories(subDept)
		if err != nil {
			return nil, err
		}
	case ToysKidsBabiesProductDepartment:
		companies, err = productDbService.GetCompaniesFromToysKidsBabies(subDept)
		if err != nil {
			return nil, err
		}
		companyCertifications, err = productDbService.GetCompanyCertificationsFromToysKidsBabies(subDept)
		if err != nil {
			return nil, err
		}
		productCertifications, err = productDbService.GetProductCertificationsFromToysKidsBabies(subDept)
		if err != nil {
			return nil, err
		}
		priceRange, err = productDbService.GetPriceRangeFromToysKidsBabies(subDept)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid department: %s", department)
	}

	// If any errors occur during database queries
	if err != nil {
		return nil, err
	}

	// Assign default values if the database returns null or empty values
	if len(companies) == 0 {
		companies = defaultCompanies
	}
	if len(companyCertifications) == 0 {
		companyCertifications = defaultCompanyCertifications
	}
	if len(productCertifications) == 0 {
		productCertifications = defaultProductCertifications
	}
	if priceRange == nil {
		priceRange = defaultPriceRange
	}

	// Construct the result
	result := &model.Filters{
		Price:                 priceRange,
		Companies:             companies,
		CompanyCertifications: companyCertifications,
		ProductCertifications: productCertifications,
	}

	return result, nil
}
