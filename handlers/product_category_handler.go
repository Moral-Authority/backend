package handlers

import (
	"errors"
	"fmt"
	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"log"
)

type ProductCategorizationService struct{}

func (s ProductService) AddCategory(request model.AddCategory, dbService database.ProductDbService) (*model.Category, error) {
	var parentID *uint
	if request.ParentID != nil && *request.ParentID != "" {
		id, err := database.StringToUint(*request.ParentID)
		if err != nil {
			return nil, err
		}
		parentID = &id
	}
	category := models.Category{
		Name:     request.Name,
		Type:     fmt.Sprintf("%s", request.Type),
		ParentID: parentID,
	}

	addedCategory, err := dbService.AddCategory(category)
	if err != nil || addedCategory == nil {
		return nil, errors.New("unable to save category to db with error %s")
	}

	log.Printf("Added category: %+v\n", addedCategory) // Log the added category

	categoryResponse := toCategoryResponse(addedCategory)
	if categoryResponse == nil {
		log.Println("toCategoryResponse returned null") // Log if toCategoryResponse returns null
	}

	return categoryResponse, nil
}

func (s ProductService) GetAllCategories(dbService database.ProductDbService) ([]*model.Category, error) {
	categories, err := dbService.GetAllCategories()
	if  err != nil || categories == nil {
		return nil, errors.New("unable to get categories from db")
	}

	return toCategoriesResponse(categories), nil
}
