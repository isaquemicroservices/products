package product

import (
	"github.com/isaqueveras/products-microservice/application/product"
)

// Server implements proto interface
type Server struct {
	product.UnimplementedProductsServer
}
