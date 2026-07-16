package repository

import (
	"context"
	"sync"

	"github.com/JiroFg/prueba-products-api/internal/domain/entities"
	"github.com/JiroFg/prueba-products-api/internal/domain/exceptions"
)

type MemoryProductRepository struct {
	mu       sync.RWMutex
	products []*entities.Product
}

func NewMemoryProductRepository(products ...*entities.Product) *MemoryProductRepository {
	return &MemoryProductRepository{
		products: products,
	}
}

func (r *MemoryProductRepository) GetProducts(ctx context.Context) ([]*entities.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	products := make([]*entities.Product, len(r.products))
	copy(products, r.products)
	return products, nil
}

func (r *MemoryProductRepository) GetProduct(ctx context.Context, id string) (*entities.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, product := range r.products {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, &exceptions.NotFoundError{
		ElementName: "Product",
		ElementId:   id,
	}
}

func (r *MemoryProductRepository) CreateProduct(ctx context.Context, product *entities.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.products = append(r.products, product)
	return nil
}

func (r *MemoryProductRepository) UpdateProduct(ctx context.Context, id string, updateProduct *entities.UpdateProduct) (*entities.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, product := range r.products {
		if product.ID == id {
			if updateProduct.Name != nil {
				product.Name = *updateProduct.Name
			}
			if updateProduct.Price != nil {
				product.Price = *updateProduct.Price
			}
			return product, nil
		}
	}
	return nil, &exceptions.NotFoundError{
		ElementName: "Product",
		ElementId:   id,
	}
}

func (r *MemoryProductRepository) DeleteProduct(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, product := range r.products {
		if product.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return &exceptions.NotFoundError{
		ElementName: "Product",
		ElementId:   id,
	}
}
