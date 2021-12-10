package product

import "time"

// Product model of one product
type Product struct {
	ID          *int64
	Name        *string
	Description *string
	Price       *float64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

// ListProducts model a list of products
type ListProducts struct {
	Data []Product
}
