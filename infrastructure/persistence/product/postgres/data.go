package postgres

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/isaqueveras/products-microservice/configuration/database"
	"github.com/isaqueveras/products-microservice/infrastructure/persistence/product"
)

// PGProduct implements methods for postgres query execution
type PGProduct struct {
	DB *database.DBTransaction
}

// ShowDetails get details of a product on database
func (pg *PGProduct) ShowDetails(ProductID *int64) (res *product.Product, err error) {
	res = new(product.Product)

	query := pg.DB.Builder.
		Select(`
			TP.id,
			TP.name,
			TP.description,
			TP.price`).
		From("t_products TP").
		Where(sq.Eq{"TP.id": ProductID}).
		Limit(1)

	// Scanning data to the struct
	if err = query.Scan(&res.ID, &res.Name, &res.Description, &res.Price); err != nil {
		return res, err
	}

	return
}
