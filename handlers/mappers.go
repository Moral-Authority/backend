package handlers

import (
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
	"strconv"
)

func toUserResponse(user models.User) *model.User {
	return &model.User{
		ID:        strconv.Itoa(int(user.ID)),
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func toFavsResponse(favs []models.Favourite) []*model.Favourite {
	var response []*model.Favourite
	for _, e := range favs {
		fav := toFavResponse(e)
		response = append(response, fav)
	}
	return response
}

func toFavResponse(fav models.Favourite) *model.Favourite {
	return &model.Favourite{
		ID:      strconv.Itoa(int(fav.ID)),
		Product: toProductResponse(fav.Product),
	}
}

func toImageResponse(image models.Image) *model.Image {
	return &model.Image{
		ID:       strconv.Itoa(int(image.ID)),
		Location: image.ImageLocation,
	}
}

func toCategoriesResponse(categories []models.Categories) []*model.Category {
	var response []*model.Category
	//for _, e := range categories {
	//	cat := toCategoryResponse(e)
	//	response = append(response, cat)
	//}
	return response
}

//func toFilterResponse(filter models.Filter) *model.Filter {
//	//return &model.Filter{
//    //    ID:    strconv.Itoa(int(filter.ID)),
//    //    : filter.Value,
//    //}
//
//	return model.Filter{}
//}

func toProductsResponse(products []models.Product) []*model.Product {
	var response []*model.Product
	for _, e := range products {
		product := toProductResponse(e)
		response = append(response, product)
	}
	return response
}

func toProductResponse(product models.Product) *model.Product {

	return &model.Product{
		ID:          strconv.Itoa(int(product.ID)),
		Title:       product.Title,
		Description: product.Description,
	}
}

func toCompanyResponse(company models.Company) *model.Company {
	return &model.Company{
		//ID:            strconv.Itoa(int(company.ID)),
		URL:         &company.Url.String,
		Description: &company.Description.String,
		IsVerified:  &company.IsVerified.Bool,
		Logo:        &company.Image.String,
	}
}

func toCertificationResponse(cert models.Certification) *model.Certification {
	return &model.Certification{
		ID:               strconv.Itoa(int(cert.ID)),
		Name:             cert.Name.String,
		Logo:             &cert.Logo.String,
		Industry:         cert.Industry.String,
		Certifier:        cert.Certifier.String,
		CertifiesCompany: &cert.CertifiesCompanies.Bool,
		CertifiesProduct: &cert.CertifiesProducts.Bool,
		CertifiesProcess: &cert.CertifiesProcesses.Bool,
		CertifierContact: &cert.CertifierContactID.String,
		Audited:          &cert.Audited.Bool,
		Auditor:          &cert.Auditor.String,
		Region:           &cert.Region.String,
		Qualifiers:       &cert.Qualifiers.String,
		Sources:          &cert.Sources.String,
	}
}
