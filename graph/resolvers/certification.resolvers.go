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

// AddCertification is the resolver for the addCertification field.
func (r *mutationResolver) AddCertification(ctx context.Context, input model.AddCertification) (*model.Certification, error) {
	cert, err := handlers.CertificationService{}.AddNewCertification(input, database.CertificationDbServiceImpl{})
	if err != nil {
		return cert, err
	}

	return cert, nil
}

// GetCertifications is the resolver for the getCertifications field.
func (r *queryResolver) GetAllCertifications(ctx context.Context) ([]*model.Certification, error) {
	certs, err := handlers.CertificationService{}.GetAllCertifications(database.CertificationDbServiceImpl{})
	if err != nil {
		return certs, err
	}

	return certs, nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) UpdateCertification(ctx context.Context, input model.UpdateCertification) (*model.Certification, error) {
	panic(fmt.Errorf("not implemented: UpdateCertification - updateCertification"))
}
