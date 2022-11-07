package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/handlers"
)

// AddStyle is the resolver for the addStyle field.
func (r *mutationResolver) AddStyle(ctx context.Context, request model.AddStyleRequest) (*model.Style, error) {
	style, err := handlers.ProductService{}.AddNewStyle(request, database.ProductDbServiceImpl{})
	if err == nil {
		return style, nil
	} else {
		return nil, err
	}
}

// AddType is the resolver for the addType field.
func (r *mutationResolver) AddType(ctx context.Context, request model.AddTypeRequest) (*model.Type, error) {
	type_, err := handlers.ProductService{}.AddNewType(request, database.ProductDbServiceImpl{})
	if err == nil {
		return type_, nil
	} else {
		return nil, err
	}
}
