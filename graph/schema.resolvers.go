package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/generated"
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

// AddFav is the resolver for the addFav field.
func (r *mutationResolver) AddFav(ctx context.Context, request model.AddUserFav) ([]*model.Favourite, error) {
	fav, err := handlers.UserService{}.AddUserFav(request, database.UserDbServiceImpl{}, database.ProductDbServiceImpl{})
	if err == nil {
		return fav, nil
	} else {
		return nil, err
	}
}

// AddDepartment is the resolver for the addDepartment field.
func (r *mutationResolver) AddDepartment(ctx context.Context, request model.AddDepartmentRequest) (*model.Department, error) {
	dept, err := handlers.ProductService{}.AddNewDepartment(request, database.ProductDbServiceImpl{})
	if err == nil {
		return dept, nil
	} else {
		return nil, err
	}
}

// AddCategory is the resolver for the addCategory field.
func (r *mutationResolver) AddCategory(ctx context.Context, request model.AddCategoryRequest) (*model.Category, error) {
	cat, err := handlers.ProductService{}.AddNewCategory(request, database.ProductDbServiceImpl{})
	if err == nil {
		return cat, nil
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

// AddStyle is the resolver for the addStyle field.
func (r *mutationResolver) AddStyle(ctx context.Context, request model.AddStyleRequest) (*model.Style, error) {
	style, err := handlers.ProductService{}.AddNewStyle(request, database.ProductDbServiceImpl{})
	if err == nil {
		return style, nil
	} else {
		return nil, err
	}
}

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

// AddCompany is the resolver for the addCompany field.
func (r *mutationResolver) AddCompany(ctx context.Context, request model.AddCompanyRequest) (*model.Company, error) {
	panic(fmt.Errorf("not implemented: AddCompany - addCompany"))
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

// GetDepartments is the resolver for the getDepartments field.
func (r *queryResolver) GetDepartments(ctx context.Context) ([]*model.Department, error) {
	departments, err := handlers.ProductService{}.GetDepartments(database.ProductDbServiceImpl{})
	if err == nil {
		return departments, nil
	} else {
		return nil, err
	}
}

// GetCompany is the resolver for the getCompany field.
func (r *queryResolver) GetCompany(ctx context.Context, id string) (*model.Company, error) {
	panic(fmt.Errorf("not implemented: GetCompany - getCompany"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
