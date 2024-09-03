package handlers

import (
	"log"
	"strconv"

	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
)

func toUserResponse(user models.User) *model.User {
	return &model.User{
		ID:    strconv.Itoa(int(user.ID)),
		Email: user.Email,
		Phone: user.Phone,
	}
}


func toFavResponse(fav *models.Favorite, product *model.Product, department ProductDepartment) *model.Favorite {
	return &model.Favorite{
		ID:      strconv.Itoa(int(fav.ID)),
		Product: toProductResponse(product, department),
	}
}

func toImageResponse(image models.Image) *model.Image {
	return &model.Image{
		ID:  strconv.Itoa(int(image.ID)),
		URL: image.Url,
	}
}

func toImagesResponse(images []models.Image) []string {
	var response []string
	for _, i := range images {
		response = append(response, i.Url)
	}
	return response
}

func toCategoriesResponse(categories []*models.Category) []*model.Category {
	var response []*model.Category
	for _, e := range categories {
		cat := toCategoryResponse(e)
		response = append(response, cat)
	}
	return response
}

func toCategoryResponse(category *models.Category) *model.Category {
	if category == nil {
		return nil
	}

	categoryResponse := &model.Category{
		ID:   strconv.Itoa(int(category.ID)),
		Type: &category.Type,
		Name: category.Name,
	}

	log.Printf("Category response: %+v\n", categoryResponse) // Log the category response

	return categoryResponse
}

//func toFilterResponse(filter models.Filter) *model.Filter {
//	//return &model.Filter{
//    //    ID:    strconv.Itoa(int(filter.ID)),
//    //    : filter.Value,
//    //}
//
//	return model.Filter{}
//}

func toPurchaseInfoResponse(purchaseInfo []models.PurchaseInfo) []*model.PurchaseInfo {
	var response []*model.PurchaseInfo
	for _, e := range purchaseInfo {
		purchase := toPurchaseInfo(e)
		response = append(response, purchase)
	}
	return response
}

func toPurchaseInfo(purchaseInfo models.PurchaseInfo) *model.PurchaseInfo {
	return &model.PurchaseInfo{
		Link:  &purchaseInfo.Url,
		Price: &purchaseInfo.Price,
	}
}

func toCompanyResponse(company *models.Company) *model.Company {
	var city, state, country *string

	if company.City.Valid {
		city = &company.City.String
	}

	if company.State.Valid {
		state = &company.State.String
	}

	if company.Country.Valid {
		country = &company.Country.String
	}

	var certs []*model.Certification
	for _, c := range company.CompanyCertifications {
		cert := &model.Certification{
			ID:                 strconv.Itoa(int(c.Certification.ID)),
			Name:               c.Certification.Name.String,
			Website:            &c.Certification.Website.String,
			Logo:               &c.Certification.Logo.String,
			Description:        &c.Certification.Description.String,
			Industry:           &c.Certification.Industry.String,
			Certifier:          &c.Certification.Certifier.String,
			CertifiesCompany:   &c.Certification.CertifiesCompany.Bool,
			CertifiesProduct:   &c.Certification.CertifiesProduct.Bool,
			CertifiesProcess:   &c.Certification.CertifiesProcess.Bool,
			CertifierContactID: &c.Certification.CertifierContactID.String,
			Audited:            &c.Certification.Audited.Bool,
			Auditor:            &c.Certification.Auditor.String,
			Region:             &c.Certification.Region.String,
			Qualifiers:         &c.Certification.Qualifiers.String,
			Sources:            &c.Certification.Sources.String,
		}
		certs = append(certs, cert)
	}

	return &model.Company{
		ID:                    strconv.Itoa(int(company.ID)),
		Name:                  company.Name,
		City:                  city,
		State:                 state,
		Country:               country,
		URL:                   &company.Url.String,
		Description:           &company.Description.String,
		IsVerified:            &company.IsVerified.Bool,
		Logo:                  &company.Image.String,
		CompanyCertifications: certs,
	}
}

func toCompaniesResponse(companies []*models.Company) []*model.Company {
	var response []*model.Company
	for _, c := range companies {
		company := toCompanyResponse(c)
		response = append(response, company)
	}
	return response
}

func toCertificationsResponse(certs []models.Certification) []*model.Certification {
	var response []*model.Certification
	for _, c := range certs {
		cert := toCertificationResponse(c)
		response = append(response, cert)
	}
	return response
}

func toCertificationResponse(cert models.Certification) *model.Certification {
	var industry, certifier, logo, certifierContactID, auditor, region, qualifiers, sources, website *string
	var certifiesCompany, certifiesProduct, certifiesProcess, audited *bool

	if cert.Industry.Valid {
		industry = &cert.Industry.String
	}
	if cert.Certifier.Valid {
		certifier = &cert.Certifier.String
	}
	if cert.Logo.Valid {
		logo = &cert.Logo.String
	}
	if cert.CertifierContactID.Valid {
		certifierContactID = &cert.CertifierContactID.String
	}
	if cert.Auditor.Valid {
		auditor = &cert.Auditor.String
	}
	if cert.Region.Valid {
		region = &cert.Region.String
	}
	if cert.Qualifiers.Valid {
		qualifiers = &cert.Qualifiers.String
	}
	if cert.Sources.Valid {
		sources = &cert.Sources.String
	}
	if cert.CertifiesCompany.Valid {
		certifiesCompany = &cert.CertifiesCompany.Bool
	}
	if cert.CertifiesProduct.Valid {
		certifiesProduct = &cert.CertifiesProduct.Bool
	}
	if cert.CertifiesProcess.Valid {
		certifiesProcess = &cert.CertifiesProcess.Bool
	}

	if cert.Audited.Valid {
		audited = &cert.Audited.Bool
	}

	if cert.Website.Valid {
		website = &cert.Website.String
	}

	return &model.Certification{
		ID:                 strconv.Itoa(int(cert.ID)),
		Name:               cert.Name.String,
		Website:            website,
		Logo:               logo,
		Industry:           industry,
		Certifier:          certifier,
		CertifiesCompany:   certifiesCompany,
		CertifiesProduct:   certifiesProduct,
		CertifiesProcess:   certifiesProcess,
		CertifierContactID: certifierContactID,
		Audited:            audited,
		Auditor:            auditor,
		Region:             region,
		Qualifiers:         qualifiers,
		Sources:            sources,
	}
}

func UintPtrToStringPtr(u *uint) *string {
	if u == nil {
		return nil
	}
	s := strconv.FormatUint(uint64(*u), 10)
	return &s
}

func ConvertStringsToColorStruct(Title, Value string) *model.Color {
	return &model.Color{
		Title: &Title,
		Value: &Value,
	}

}

func FormatStringListToColorStructList(colors map[string]string) []*model.Color {
	result := []*model.Color{}

	for title, value := range colors {
		color := ConvertStringsToColorStruct(title, value)
		result = append(result, color)
	}

	return result
}

func ConvertString(s string) *string {
	return &s
}
