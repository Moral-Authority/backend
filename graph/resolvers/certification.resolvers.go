package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/handlers"
)

// AddCertification is the resolver for the addCertification field.
func (r *mutationResolver) AddCertification(ctx context.Context, input model.AddCertification) (*model.Certification, error) {
	cert, err := handlers.CertificationService{}.AddNewCertification(input, nil)
	if err != nil {
		return cert, err
	}

	return cert, nil
}

// UpdateCertification is the resolver for the updateCertification field.
func (r *mutationResolver) UpdateCertification(ctx context.Context, input model.UpdateCertification) (*model.Certification, error) {
	certs, err := handlers.CertificationService{}.UpdateCertification(input, nil)
	if err != nil {
		return certs, err
	}

	return nil, nil
}

// GetCertifications is the resolver for the getCertifications field.
func (r *queryResolver) GetAllCertifications(ctx context.Context) ([]*model.Certification, error) {
	certs, err := handlers.CertificationService{}.GetAllCertifications(nil)
	if err != nil {
		return certs, err
	}

	return certs, nil
}
