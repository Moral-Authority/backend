package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/howstrongiam/backend/graph/model"
)

// AddCategory is the resolver for the addCategory field.
func (r *mutationResolver) AddCategory(ctx context.Context, input model.AddCategory) (*model.Section, error) {
	//cat, err := handlers.ProductService{}.AddNewCategory(input, database.ProductDbServiceImpl{})
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
}

// GetAllCategories is the resolver for the getAllCategories field.
func (r *queryResolver) GetAllCategories(ctx context.Context) ([]*model.Section, error) {
	//type_, err := handlers.ProductService{}.AddNewType(input, database.ProductDbServiceImpl{})
	//if err == nil {
	//	return type_, nil
	//} else {
	//	return nil, err
	//}
	panic(fmt.Errorf("not implemented: GetAllCategories - getAllCategories"))
}
