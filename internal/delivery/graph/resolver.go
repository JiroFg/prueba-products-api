package graph

import "github.com/JiroFg/prueba-products-api/internal/usecases"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	ProductUseCase *usecases.ProductUseCase
}
