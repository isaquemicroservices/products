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

// List list all products on database
func (s *Server) List(ctx context.Context, _ *product.Void) (res *product.ListProducts, err error) {
	if res, err = product.ListAll(ctx); err != nil {
		return nil, err
	}

	return
}

// Add create product on database
func (s *Server) Add(ctx context.Context, in *product.Product) (*product.Void, error) {
	if err := product.Add(ctx, in); err != nil {
		return nil, err
	}

	return &product.Void{}, nil
}

// Add create product on database
func (s *Server) ListAllProductsWithMinimumQuantity(ctx context.Context, _ *product.Void) (*product.ListProducts, error) {
	if err := product.ListAllProductsWithMinimumQuantity(ctx); err != nil {
		return nil, err
	}

	return &product.ListProducts{}, nil
}
