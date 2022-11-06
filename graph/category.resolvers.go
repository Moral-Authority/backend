package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/handlers"

	"github.com/howstrongiam/backend/graph/model"
)

// AddCategory is the resolver for the addCategory field.
func (r *mutationResolver) AddCategory(ctx context.Context, request model.AddCategoryRequest) (*model.Category, error) {
	cat, err := handlers.ProductService{}.AddNewCategory(request, database.ProductDbServiceImpl{})
	if err == nil {
		return cat, nil
	} else {
		return nil, err
	}
}
