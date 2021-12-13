package product

import (
	"context"

	"github.com/isaqueveras/products-microservice/configuration/database"
	"github.com/isaqueveras/products-microservice/domain/product"
	infra "github.com/isaqueveras/products-microservice/infrastructure/persistence/product"
	"github.com/isaqueveras/products-microservice/utils"
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

	// making a list to products
	res.Products = make([]*Product, len(data.Data))
	for i := range data.Data {
		res.Products[i] = &Product{
			Id:          *data.Data[i].ID,
			Name:        *data.Data[i].Name,
			Description: *data.Data[i].Description,
			Price:       *data.Data[i].Price,
		}
	}

	return
}

// Add add a product on database
func Add(ctx context.Context, in *Product) (err error) {
	var (
		transaction *database.DBTransaction
		data        = infra.Product{
			ID:          utils.GetPointerInt64(in.GetId()),
			Name:        utils.GetPointerString(in.GetName()),
			Description: utils.GetPointerString(in.GetDescription()),
			Price:       utils.GetPointerFloat64(in.GetPrice()),
		}
	)

	// opening connection with database
	if transaction, err = database.OpenConnection(ctx, false); err != nil {
		return err
	}

	// rollback on transaction
	defer transaction.Rollback()

	// initialize repository of product
	var repo = product.GetProductRepository(transaction)
	if err = repo.Add(&data); err != nil {
		return err
	}

	// commit on transaction
	if err = transaction.Commit(); err != nil {
		return err
	}

	return
}
