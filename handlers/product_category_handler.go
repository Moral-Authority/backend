package handlers

import (
	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
)

type ProductCategorizationService struct{}

func (s ProductService) AddCategoryHandler(request model.AddCategory, dbService database.ProductDbService) (*model.Category, error) {
	// var parentID *uint
	// if request.ParentID != nil && *request.ParentID != "" {
	// 	id, err := database.StringToUint(*request.ParentID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	parentID = &id
	// }
	// category := models.Category{
	// 	Name:     request.Name,
	// 	Type:     fmt.Sprintf("%s", request.Type),
	// 	// ParentID: parentID,
	// }

	// addedCategory, err := dbService.AddCategory(category)
	// if err != nil || addedCategory == nil {
	// 	return nil, errors.New("unable to save category to db with error %s")
	// }

	// log.Printf("Added category: %+v\n", addedCategory) // Log the added category

	// categoryResponse := toCategoryResponse(addedCategory)
	// if categoryResponse == nil {
	// 	log.Println("toCategoryResponse returned null") // Log if toCategoryResponse returns null
	// }

	return nil, nil
}

func (s ProductService) GetAllCategoriesHandler(dbService database.ProductDbService) ([]*model.Category, error) {
	// categories, err := dbService.GetAllCategories()
	// if  err != nil || categories == nil {
	// 	return nil, errors.New("unable to get categories from db")
	// }

	return nil, nil
}
