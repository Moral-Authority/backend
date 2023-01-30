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

// AddSection is the resolver for the addSection field.
func (r *mutationResolver) AddSection(ctx context.Context, input model.AddSection) (*model.Section, error) {
	dept, err := handlers.ProductService{}.AddNewSection(input, database.ProductDbServiceImpl{})
	if err == nil {
		return dept, nil
	} else {
		return nil, err
	}
}

// AddSubSection is the resolver for the addSubSection field.
func (r *mutationResolver) AddSubSection(ctx context.Context, input model.AddSubSection) (*model.SubSection, error) {
	dept, err := handlers.ProductService{}.AddNewSubSection(input, database.ProductDbServiceImpl{})
	if err == nil {
		return dept, nil
	} else {
		return nil, err
	}
}

// AddDepartment is the resolver for the addDepartment field.
func (r *mutationResolver) AddDepartment(ctx context.Context, input model.AddDepartment) (*model.Department, error) {
	dept, err := handlers.ProductService{}.AddNewDepartment(input, database.ProductDbServiceImpl{})
	if err == nil {
		return dept, nil
	} else {
		return nil, err
	}
}

// AddCategory is the resolver for the addCategory field.
func (r *mutationResolver) AddCategory(ctx context.Context, input model.AddCategory) (*model.Category, error) {
	cat, err := handlers.ProductService{}.AddNewCategory(input, database.ProductDbServiceImpl{})
	if err == nil {
		return cat, nil
	} else {
		return nil, err
	}
}

// AddSubCategory is the resolver for the addSubCategory field.
func (r *mutationResolver) AddSubCategory(ctx context.Context, input model.AddSubCategory) (*model.SubCategory, error) {
	dept, err := handlers.ProductService{}.AddNewSubCategory(input, database.ProductDbServiceImpl{})
	if err == nil {
		return dept, nil
	} else {
		return nil, err
	}
}

// AddType is the resolver for the addType field.
func (r *mutationResolver) AddType(ctx context.Context, input model.AddTypeRequest) (*model.Type, error) {
	type_, err := handlers.ProductService{}.AddNewType(input, database.ProductDbServiceImpl{})
	if err == nil {
		return type_, nil
	} else {
		return nil, err
	}
}

// AddStyle is the resolver for the addStyle field.
func (r *mutationResolver) AddStyle(ctx context.Context, input model.AddStyleRequest) (*model.Style, error) {
	style, err := handlers.ProductService{}.AddNewStyle(input, database.ProductDbServiceImpl{})
	if err == nil {
		return style, nil
	} else {
		return nil, err
	}
}

// AddProductFilter is the resolver for the addProductFilter field.
func (r *mutationResolver) AddProductFilter(ctx context.Context, input model.AddFilter) (string, error) {
	style, err := handlers.ProductService{}.AddNewFilter(input, database.ProductDbServiceImpl{})
	if err == nil {
		return style, nil
	} else {
		return nil, err
	}
}

// GetSections is the resolver for the getSections field.
func (r *queryResolver) GetSections(ctx context.Context, input *string) ([]*model.Section, error) {
	panic(fmt.Errorf("not implemented: GetSections - getSections"))
}

// GetSubSections is the resolver for the getSubSections field.
func (r *queryResolver) GetSubSections(ctx context.Context, input *string) ([]*model.SubSection, error) {
	panic(fmt.Errorf("not implemented: GetSubSections - getSubSections"))
}

// GetDepartments is the resolver for the getDepartments field.
func (r *queryResolver) GetDepartments(ctx context.Context, input *string) ([]*model.Department, error) {
	panic(fmt.Errorf("not implemented: GetDepartments - getDepartments"))
}

// GetCategories is the resolver for the getCategories field.
func (r *queryResolver) GetCategories(ctx context.Context, input *string) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented: GetCategories - getCategories"))
}

// GetSubCategories is the resolver for the getSubCategories field.
func (r *queryResolver) GetSubCategories(ctx context.Context, input *string) ([]*model.SubCategory, error) {
	panic(fmt.Errorf("not implemented: GetSubCategories - getSubCategories"))
}

// GetTypes is the resolver for the getTypes field.
func (r *queryResolver) GetTypes(ctx context.Context, input *string) ([]*model.Type, error) {
	panic(fmt.Errorf("not implemented: GetTypes - getTypes"))
}

// GetStyle is the resolver for the getStyle field.
func (r *queryResolver) GetStyle(ctx context.Context, input *string) ([]*model.Style, error) {
	panic(fmt.Errorf("not implemented: GetStyle - getStyle"))
}

// GetFilters is the resolver for the getFilters field.
func (r *queryResolver) GetFilters(ctx context.Context, input *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: GetFilters - getFilters"))
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) AddProductOrganizationType(ctx context.Context, input model.CategorizationInput) (*model.ProductOrganizationType, error) {
	panic(fmt.Errorf("not implemented: AddProductOrganizationType - addProductOrganizationType"))
}
