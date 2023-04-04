package handlers

import (
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
)

type ProductCategorizationService struct{}

func (s ProductService) AddCategory(request model.AddCategory, dbService database.ProductDbService) (*model.Category, error) {
	//parentID, err := database.StringToUint(request.ParentID)
	//if err != nil {
	//	return nil, err
	//}
	//category := models.Categories{
	//	Name:     request.Name,
	//	ParentId: parentID,
	//}
	//addedCategory := dbService.AddCategory(category)
	//if addedCategory == nil {
	//	return nil, errors.New("unable to save category to db")
	//}
	//return toCategoryResponse(*addedCategory), nil
	return nil, nil
}
func (s ProductService) GetAllCateogories(dbService database.ProductDbService) ([]*model.Category, error) {
	//sections := dbService.GetAllCategories()
	//if sections == nil {
	//	return nil, errors.New("unable to get departments from db")
	//}
	//return sections, nil
	//return toSectionResponse(sections), nil
	return nil, nil
}
