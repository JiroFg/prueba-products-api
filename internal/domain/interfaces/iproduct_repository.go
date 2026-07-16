package interfaces

import (
	"context"

	"github.com/JiroFg/prueba-products-api/internal/domain/entities"
)

type IProductRepository interface {
	GetProducts(ctx context.Context) ([]*entities.Product, error)
	GetProduct(ctx context.Context, id string) (*entities.Product, error)
	CreateProduct(ctx context.Context, product *entities.Product) error
	UpdateProduct(ctx context.Context, id string, updateProduct *entities.UpdateProduct) error
	DeleteProduct(ctx context.Context, id string) error
}
