package handlers

import (
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
)

func buildFiltersMap(filter *model.ProductFilterInput) map[string]interface{} {

	filters := make(map[string]interface{})

	if filter == nil {
		return filters
	}

	// Add price range to the filter map if present
	if filter.PriceRange != nil {
		filters["priceMin"] = filter.PriceRange.Min
		filters["priceMax"] = filter.PriceRange.Max
	}

	// Dereference company certifications pointers
	if len(filter.CompanyCertifications) > 0 {
		certifications := make([]string, len(filter.CompanyCertifications))
		for i, cert := range filter.CompanyCertifications {
			if cert != nil {
				certifications[i] = *cert
			}
		}
		filters["companyCertifications"] = certifications
	}

	// Dereference product certifications pointers
	if len(filter.ProductCertifications) > 0 {
		certifications := make([]string, len(filter.ProductCertifications))
		for i, cert := range filter.ProductCertifications {
			if cert != nil {
				certifications[i] = *cert
			}
		}
		filters["productCertifications"] = certifications
	}

	// Dereference companies pointers
	if len(filter.Companies) > 0 {
		companies := make([]string, len(filter.Companies))
		for i, company := range filter.Companies {
			if company != nil {
				companies[i] = *company
			}
		}
		filters["companies"] = companies
	}

	return filters
}

func FetchProductAndCompanyCertsConcurrently(
	productID uint,
	getCompanyCertsFunc func(uint) ([]*models.Certification, error),
	getProductCertsFunc func(uint) ([]*models.Certification, error),
) ([]*models.Certification, []*models.Certification, error) {

	companyCertsChan := make(chan []*models.Certification)
	productCertsChan := make(chan []*models.Certification)
	errChan := make(chan error)

	go func() {
		companyCerts, err := getCompanyCertsFunc(productID)
		if err != nil {
			errChan <- err
			return
		}
		companyCertsChan <- companyCerts
	}()

	go func() {
		productCerts, err := getProductCertsFunc(productID)
		if err != nil {
			errChan <- err
			return
		}
		productCertsChan <- productCerts
	}()

	// Wait for results and errors

	var companyCerts []*models.Certification
	var productCerts []*models.Certification

	for i := 0; i < 2; i++ {
		select {
		case err := <-errChan:
			return nil, nil, err
		case cc := <-companyCertsChan:
			companyCerts = cc
		case pc := <-productCertsChan:
			productCerts = pc
		}
	}

	// Combine data into a new response
	return productCerts, companyCerts, nil
}

// func AddProductByDepartment(
// 	department ProductDepartment,
// 	product interface{},
// 	request model.AddProductRequest,
// 	productDbService database.ProductDbService,
// 	imageDbService database.ImageDbService,
// 	certificationService database.CertificationDbService,
// ) (*model.Product, error) {

// 	addedProduct, err := productDbService.AddProduct(product)
// 	if err != nil || addedProduct == nil {
// 		return nil, errors.New("unable to save product to db")
// 	}

// 	// Type assertion based on department to access the ID
// 	var productID uint
// 	switch department {
// 	case HomeGardenProductDepartment:
// 		productID = addedProduct.(*models.HomeGardenProduct).ID
// 	case BathBeautyProductDepartment:
// 		productID = addedProduct.(*models.HealthBathBeautyProduct).ID
// 	case ClothingAccessoriesProductDepartment:
// 		productID = addedProduct.(*models.ClothingAccessoriesProduct).ID
// 	case ToysKidsBabiesProductDepartment:
// 		productID = addedProduct.(*models.ToysKidsBabiesProduct).ID
// 	default:
// 		return nil, fmt.Errorf("unknown department type: %d", department)
// 	}

// 	// Handle image links
// 	for _, i := range request.ImageLinks {
// 		image := models.Image{
// 			Url: *i,
// 		}

// 		addedImage, err := imageDbService.AddImage(image)
// 		if err != nil || addedImage == nil {
// 			return nil, errors.New("unable to save image to db")
// 		}
// 	}

// 	// Handle certifications
// 	for _, c := range request.Certifications {
// 		foundCert, err := certificationService.GetCertificationById(*c.CertificationID)
// 		if err != nil {
// 			return nil, fmt.Errorf("unable to find certification number %d", c.CertificationID)
// 		}

// 		cert := models.ProductCertification{
// 			ProductID:       productID,
// 			CertificationID: foundCert.ID,
// 		}

// 		addedProductCert, err := productDbService.AddProductCertification(cert)
// 		if err != nil || addedProductCert == nil {
// 			return nil, errors.New("unable to save product certification to db")
// 		}
// 	}

// 	// Convert the saved product to the common response type model.Product
// 	return toProductResponse(addedProduct, department), nil
// }
