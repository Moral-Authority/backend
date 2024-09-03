package handlers

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
