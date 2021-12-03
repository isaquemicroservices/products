package product

import (
	"context"

	"github.com/isaqueveras/products-microservice/application/product"
)

// Server implements proto interface
type Server struct {
	product.UnimplementedProductsServer
}

// Show get details of a product
func (s *Server) Show(ctx context.Context, in *product.Params) (res *product.Product, err error) {
	if res, err = product.ShowDetails(ctx, in); err != nil {
		return nil, err
	}

	return
}
