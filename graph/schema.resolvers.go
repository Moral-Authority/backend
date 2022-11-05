package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/generated"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/handlers"
)

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := handlers.UserService{}.AddNewUser(model.NewUser{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}, database.UserDbServiceImpl{})
	if err == nil {
		return user, nil
	} else {
		return nil, err
	}
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	user, err := handlers.UserService{}.UpdateUser(model.UpdateUser{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}, database.UserDbServiceImpl{})
	if err == nil {
		return user, nil
	} else {
		return nil, err
	}
}

// AddFav is the resolver for the addFav field.
func (r *mutationResolver) AddFav(ctx context.Context, request model.AddUserFav) ([]*model.Favourite, error) {
	fav, err := handlers.UserService{}.AddUserFav(model.AddUserFav{
		UserID:    request.UserID,
		ProductID: request.ProductID,
	}, database.UserDbServiceImpl{}, database.ProductDbServiceImpl{})
	if err == nil {
		return fav, nil
	} else {
		return nil, err
	}
}

// AddDepartment is the resolver for the addDepartment field.
func (r *mutationResolver) AddDepartment(ctx context.Context, request model.AddDepartmentRequest) (*model.Department, error) {
	dept, err := handlers.ProductService{}.AddNewDepartment(model.AddDepartmentRequest{
		Title: request.Title,
	})
	if err == nil {
		return dept, nil
	} else {
		return nil, err
	}
}

// AddType is the resolver for the addType field.
func (r *mutationResolver) AddType(ctx context.Context, request model.AddTypeRequest) (*model.Type, error) {
	type_, err := handlers.ProductService{}.AddNewType(model.AddTypeRequest{
		Title:      request.Title,
		CategoryID: request.CategoryID,
	})
	if err == nil {
		return type_, nil
	} else {
		return nil, err
	}
}

// AddStyle is the resolver for the addStyle field.
func (r *mutationResolver) AddStyle(ctx context.Context, request model.AddStyleRequest) (*model.Style, error) {
	style, err := handlers.ProductService{}.AddNewStyle(model.AddStyleRequest{
		Title:  request.Title,
		TypeID: request.TypeID,
	})
	if err == nil {
		return style, nil
	} else {
		return nil, err
	}
}

// AddProduct is the resolver for the addProduct field.
func (r *mutationResolver) AddProduct(ctx context.Context, request model.AddProductRequest) (*model.Product, error) {
	product, err := handlers.ProductService{}.AddNewProduct(model.AddProductRequest{
		Title:         request.Title,
		URL:           request.URL,
		Description:   request.Description,
		UserID:        request.UserID,
		ImageLocation: request.ImageLocation,
		Certification: request.Certification,
		StyleID:       request.StyleID,
	})
	if err == nil {
		return product, nil
	} else {
		return nil, err
	}
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := handlers.UserService{}.GetUserById(id)
	if err == nil {
		return user, nil
	} else {
		return nil, err
	}
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := handlers.UserService{}.GetUsers()
	if err == nil {
		return users, nil
	} else {
		return nil, err
	}
}

// GetDepartments is the resolver for the getDepartments field.
func (r *queryResolver) GetDepartments(ctx context.Context) ([]*model.Department, error) {
	departments, err := handlers.ProductService{}.GetDepartments()
	if err == nil {
		return departments, nil
	} else {
		return nil, err
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
