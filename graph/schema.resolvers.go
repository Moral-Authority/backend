package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/howstrongiam/backend/graph/generated"
	"github.com/howstrongiam/backend/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	createdAt := time.Now().Format("2006-01-02")
	
	user := model.User{
		Name: input.Name,
		CreatedAt: createdAt,
		Age: input.Age,
		Address: input.Address,
	}

	_, err := r.DB.Model(&user).Insert()
   if err != nil {
       return nil, fmt.Errorf("error inserting new user: %v", err)
   }

   return &user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	err := r.DB.Model(&users).Select()
	if err != nil {
		return nil, fmt.Errorf("error retrieve users: %v", err)
	}

	return users, nil
}


// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

