package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/handlers"
)

// AddDepartment is the resolver for the addDepartment field.
func (r *mutationResolver) AddDepartment(ctx context.Context, request model.AddDepartmentRequest) (*model.Department, error) {
	dept, err := handlers.ProductService{}.AddNewDepartment(request, database.ProductDbServiceImpl{})
	if err == nil {
		return dept, nil
	} else {
		return nil, err
	}
}

// GetDepartments is the resolver for the getDepartments field.
func (r *queryResolver) GetDepartments(ctx context.Context) ([]*model.Department, error) {
	panic(fmt.Errorf("not implemented: GetDepartments - getDepartments"))
}
