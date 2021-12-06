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

	if data != nil {
		res.Id = *data.ID
		res.Name = *data.Name
		res.Description = *data.Description
		res.Price = *data.Price
	}

	return
}

// ListAll get a list of products on database
func ListAll(ctx context.Context) (res *ListProducts, err error) {
	res = new(ListProducts)
	var transaction *database.DBTransaction

	// opening connection with database
	if transaction, err = database.OpenConnection(ctx, true); err != nil {
		return res, err
	}

	// rollback on transaction
	defer transaction.Rollback()

	// initialize repository of product
	var repo = product.GetProductRepository(transaction)

	// get a list of a product on database
	data, err := repo.ListAll()
	if err != nil {
		return res, err
	}

	res.Products = make([]Product, len(data.Data))
	for ii := range res.Products {
		if data != nil {
			res.Products[ii].Id = *data.Data[ii].ID
			res.Products[ii].Name = *data.Data[ii].Name
			res.Products[ii].Description = *data.Data[ii].Description
			res.Products[ii].Price = *data.Data[ii].Price
		}
	}

	return
}
