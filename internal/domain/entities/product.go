package entities

type Product struct {
	ID        string
	Name      string
	Price     float64
	Stock     int32
	CreatedAt string
}

type UpdateProduct struct {
	Name  *string
	Price *float64
}
