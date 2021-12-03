package product

import (
	"github.com/isaqueveras/products-microservice/configuration/database"
	"github.com/isaqueveras/products-microservice/infrastructure/persistence/product"
	"github.com/isaqueveras/products-microservice/infrastructure/persistence/product/postgres"
)

// repository is a base structure that implements methods specified by IProduct
type repository struct {
	pg *postgres.PGProduct
}

// New creates a new product repository
func New(db *database.DBTransaction) *repository {
	return &repository{pg: &postgres.PGProduct{DB: db}}
}

// ShowDetails get details of a product on database
func (r *repository) ShowDetails(ProductID *int64) (*product.Product, error) {
	return r.pg.ShowDetails(ProductID)
}
