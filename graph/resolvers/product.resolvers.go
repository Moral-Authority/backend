package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/handlers"
)

// AddProduct is the resolver for the addProduct field.
func (r *mutationResolver) AddProduct(ctx context.Context, request model.AddProductRequest) (*model.Product, error) {
	product, err := handlers.ProductService{}.AddNewProduct(
		request,
		database.ProductDbServiceImpl{},
		database.ImageDbServiceImpl{},
		database.CertificatesDbServiceImpl{},
	)
	if err == nil {
		return product, nil
	} else {
		return nil, err
	}
}
