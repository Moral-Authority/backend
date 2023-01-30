package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/handlers"
)

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := handlers.UserService{}.AddNewUser(input, database.UserDbServiceImpl{})
	if err == nil {
		return user, nil
	} else {
		return nil, err
	}
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	user, err := handlers.UserService{}.UpdateUser(input, database.UserDbServiceImpl{})
	if err == nil {
		return user, nil
	} else {
		return nil, err
	}
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := handlers.UserService{}.GetUserById(id, database.UserDbServiceImpl{})
	if err == nil {
		return user, nil
	} else {
		return nil, err
	}
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := handlers.UserService{}.GetUsers(database.UserDbServiceImpl{})
	if err == nil {
		return users, nil
	} else {
		return nil, err
	}
}
