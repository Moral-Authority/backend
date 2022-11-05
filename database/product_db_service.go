package database

import "github.com/howstrongiam/backend/models"

type ProductDbService interface {
	GetProduct(string) *models.Product
}
