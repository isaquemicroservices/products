package product

// Product model of one product
type Product struct {
	ID          *int64
	Name        *string
	Description *string
	Price       *float64
}
