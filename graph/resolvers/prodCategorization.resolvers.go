package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/handlers"
)

// AddCategory is the resolver for the addCategory field.
func (r *mutationResolver) AddCategory(ctx context.Context, input model.AddCategory) (*model.Category, error) {
	cat, err := handlers.ProductService{}.AddCategory(input, database.ProductDbServiceImpl{})
	if err != nil {
		return nil, err
	}

	return cat, nil
}

// GetAllCategories is the resolver for the getAllCategories field.
func (r *queryResolver) GetAllCategories(ctx context.Context) ([]*model.Category, error) {
	categories, err := handlers.ProductService{}.GetAllCategories(database.ProductDbServiceImpl{})
	if err != nil {
		return categories, nil
	}
	return categories, nil
}
