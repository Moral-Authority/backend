package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/Moral-Authority/backend/graph/generated"
)

// BaseMutation is the resolver for the BaseMutation field.
func (r *mutationResolver) BaseMutation(ctx context.Context) (any, error) {
	panic(fmt.Errorf("not implemented: BaseMutation - BaseMutation"))
}

// BaseQuery is the resolver for the BaseQuery field.
func (r *queryResolver) BaseQuery(ctx context.Context) (any, error) {
	panic(fmt.Errorf("not implemented: BaseQuery - BaseQuery"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
