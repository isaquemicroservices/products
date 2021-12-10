package postgres

import (
	"database/sql"

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

// ListAll get a list of a products on database
func (pg *PGProduct) ListAll() (res *product.ListProducts, err error) {
	res = new(product.ListProducts)

	var (
		product product.Product
		lines   *sql.Rows
	)

	if lines, err = pg.DB.Builder.
		Select(`
			TP.id,
			TP.name,
			TP.description,
			TP.price`).
		From("t_products TP").
		Query(); err != nil {
		return res, err
	}

	for lines.Next() {
		// scanning data to the struct
		if err = lines.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			return res, err
		}

		// append products on list
		res.Data = append(res.Data, product)
	}

	return
}

// Add create a products on database
func (pg *PGProduct) Add(in *product.Product) (err error) {
	if err = pg.DB.Builder.
		Insert("t_products").
		Columns("name", "description", "price").
		Values(in.Name, in.Description, in.Price).
		Suffix("RETURNING id").
		QueryRow().
		Scan(new(int64)); err != nil {
		return err
	}

	return nil
}
