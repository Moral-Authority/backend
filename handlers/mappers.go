package handlers

import (
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
	"strconv"
)

func toUserResponse(user models.User) *model.User {
	return &model.User{
		ID:         strconv.Itoa(int(user.ID)),
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Favourites: toFavsResponse(user.Favourites),
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

func toDepartmentsResponse(departments []*models.Department) []*model.Department {
	var response []*model.Department
	for _, e := range departments {
		dept := toDepartmentResponse(*e)
		response = append(response, dept)
	}
	return response
}

func toDepartmentResponse(department models.Department) *model.Department {
	return &model.Department{
		ID:         strconv.Itoa(int(department.ID)),
		Title:      department.Title,
		Categories: toCategoriesResponse(department.Categories),
	}
}

func toCategoriesResponse(categories []models.Category) []*model.Category {
	var response []*model.Category
	for _, e := range categories {
		cat := toCategoryResponse(e)
		response = append(response, cat)
	}
	return response
}

func toCategoryResponse(category models.Category) *model.Category {
	return &model.Category{
		ID:    strconv.Itoa(int(category.ID)),
		Title: category.Title,
		Types: toTypesResponse(category.Types),
	}
}

func toTypesResponse(types []models.Type) []*model.Type {
	var response []*model.Type
	for _, e := range types {
		type_ := toTypeResponse(e)
		response = append(response, type_)
	}
	return response
}

func toTypeResponse(type_ models.Type) *model.Type {
	return &model.Type{
		ID:     strconv.Itoa(int(type_.ID)),
		Title:  type_.Title,
		Styles: toStylesResponse(type_.Styles),
	}
}

func toStylesResponse(styles []models.Style) []*model.Style {
	var response []*model.Style
	for _, e := range styles {
		style := toStyleResponse(e)
		response = append(response, style)
	}
	return response
}

func toStyleResponse(style models.Style) *model.Style {
	return &model.Style{
		ID:       strconv.Itoa(int(style.ID)),
		Products: toProductsResponse(style.Products),
	}
}

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
		ID:            strconv.Itoa(int(product.ID)),
		Title:         product.Title,
		URL:           product.Url,
		Description:   product.Description,
		UserID:        strconv.Itoa(int(product.UserId)),
		Image:         toImageResponse(product.Image),
		Certification: toCertificationResponse(product.Certification),
	}
}

func toCompanyResponse(company models.Company) *model.Company {
	return &model.Company{
		ID:            strconv.Itoa(int(company.ID)),
		URL:           company.Url,
		Description:   company.Description,
		User:          toUserResponse(company.User),
		IsVerified:    company.IsVerified,
		Image:         toImageResponse(company.Image),
		Certification: toCertificationResponse(company.Certification),
	}
}

func toCertificationResponse(cert models.Certification) *model.Certification {
	return &model.Certification{
		ID:                strconv.Itoa(int(cert.ID)),
		CertifyingCompany: cert.CertifyingCompany,
		CertName:          cert.CertName,
	}
}
