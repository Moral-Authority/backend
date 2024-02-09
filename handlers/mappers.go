package handlers

import (
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"log"
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
		ID: strconv.Itoa(int(fav.ID)),
		//Product: toProductResponse(fav.Product),
	}
}

func toImageResponse(image models.Image) *model.Image {
	return &model.Image{
		ID:       strconv.Itoa(int(image.ID)),
		Location: image.ImageLocation,
	}
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
		ID:       strconv.Itoa(int(category.ID)),
		Type:     &category.Type,
		Name:     category.Name,
		ParentID: UintPtrToStringPtr(category.ParentID),
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
		//Certification: toCertificationResponse(product.Certification),
	}
}

func toCompanyResponse(company *models.Company) *model.Company {
	return &model.Company{
		//ID:            strconv.Itoa(int(company.ID)),
		URL:         &company.Url.String,
		Description: &company.Description.String,
		IsVerified:  &company.IsVerified.Bool,
		Logo:        &company.Image.String,
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

func toCertificationResponse(cert models.Certification) *model.Certification {
    var industry, certifier, logo, certifierContactID, auditor, region, qualifiers, sources *string
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

    return &model.Certification{
        ID:                strconv.Itoa(int(cert.ID)),
        Name:              cert.Name.String,
        Logo:              logo,
        Industry:          industry,
        Certifier:         certifier,
        CertifiesCompany:  certifiesCompany,
        CertifiesProduct:  certifiesProduct,
        CertifiesProcess:  certifiesProcess,
        CertifierContactID :certifierContactID,
        Audited:           audited,
        Auditor:           auditor,
        Region:            region,
        Qualifiers:        qualifiers,
        Sources:           sources,
    }
}

func UintPtrToStringPtr(u *uint) *string {
	if u == nil {
		return nil
	}
	s := strconv.FormatUint(uint64(*u), 10)
	return &s
}
