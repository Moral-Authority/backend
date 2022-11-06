package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/handlers"

	"github.com/howstrongiam/backend/graph/model"
)

// AddFav is the resolver for the addFav field.
func (r *mutationResolver) AddFav(ctx context.Context, request model.AddUserFav) ([]*model.Favourite, error) {
	fav, err := handlers.UserService{}.AddUserFav(request, database.UserDbServiceImpl{}, database.ProductDbServiceImpl{})
	if err == nil {
		return fav, nil
	} else {
		return nil, err
	}
}
