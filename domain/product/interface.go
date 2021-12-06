package product

import (
	"github.com/isaqueveras/products-microservice/infrastructure/persistence/product"
)

// IProduct defines all services available for product
type IProduct interface {
	ShowDetails(*int64) (*product.Product, error)
	ListAll() (*product.ListProducts, error)
}
