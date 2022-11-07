package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/howstrongiam/backend/graph/model"
)

// AddCompany is the resolver for the addCompany field.
func (r *mutationResolver) AddCompany(ctx context.Context, request model.AddCompanyRequest) (*model.Company, error) {
	panic(fmt.Errorf("not implemented: AddCompany - addCompany"))
}

// GetCompany is the resolver for the getCompany field.
func (r *queryResolver) GetCompany(ctx context.Context, id string) (*model.Company, error) {
	panic(fmt.Errorf("not implemented: GetCompany - getCompany"))
}
