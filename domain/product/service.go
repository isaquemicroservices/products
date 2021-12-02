package product

import "github.com/isaqueveras/products-microservice/configuration/database"

// Service models a service base struct
type Service struct {
	repo IProduct
}

// GetService retrieves a service type
func GetService(r IProduct) *Service {
	return &Service{repo: r}
}

// GetProductRepository retrieve repository for access to product data
func GetProductRepository(db *database.DBTransaction) IProduct {
	return New(db)
}
