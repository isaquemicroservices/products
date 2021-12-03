package product

import (
	"context"
	"log"
	"testing"

	app "github.com/isaqueveras/products-microservice/application/product"
	"github.com/isaqueveras/products-microservice/configuration/database"
	"gopkg.in/go-playground/assert.v1"
)

var server app.ProductsServer = &Server{}

func TestHandler(t *testing.T) {
	t.Run("TestShowDetails", func(t *testing.T) {
		productID := getValidProduct()
		_, err := server.Show(context.Background(), &app.Params{Id: *productID})
		assert.Equal(t, err, nil)
	})
}

func getValidProduct() (id *int64) {
	transaction, err := database.OpenConnection(context.Background(), true)
	if err != nil {
		log.Fatal(err)
	}

	defer transaction.Rollback()

	if err = transaction.Builder.
		Select("TP.id").
		From("t_products TP").
		Limit(1).
		Scan(&id); err != nil {
		log.Fatal(err)
	}

	return
}
