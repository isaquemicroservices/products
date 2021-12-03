package product

import (
	"context"

	"github.com/isaqueveras/products-microservice/configuration/database"
	"github.com/isaqueveras/products-microservice/domain/product"
)

// ShowDetails show details of a product
func ShowDetails(ctx context.Context, in *Params) (res *Product, err error) {
	res = new(Product)

	var (
		transaction *database.DBTransaction
		productID   = in.GetId()
	)

	// Opening connection with database
	if transaction, err = database.OpenConnection(ctx, true); err != nil {
		return res, err
	}

	// Rollback on transaction
	defer transaction.Rollback()

	// Initialize repository of product
	var repo = product.GetProductRepository(transaction)

	// Get details of a product on database
	data, err := repo.ShowDetails(&productID)
	if err != nil {
		return res, err
	}

	res.Id = *data.ID
	res.Name = *data.Name
	res.Description = *data.Description
	res.Price = *data.Price

	return
}
