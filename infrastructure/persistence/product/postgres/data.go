package postgres

import "github.com/isaqueveras/products-microservice/configuration/database"

// PGProduct implements methods for postgres query execution
type PGProduct struct {
	DB *database.DBTransaction
}
