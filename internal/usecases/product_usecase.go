package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/JiroFg/prueba-products-api/internal/domain/entities"
	"github.com/JiroFg/prueba-products-api/internal/domain/interfaces"
)

type ProductUseCase struct {
	repo interfaces.IProductRepository
}

func NewProductUseCase(repo interfaces.IProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repo: repo,
	}
}

func (p *ProductUseCase) GetProducts(ctx context.Context) ([]*entities.Product, error) {
	products, err := p.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductUseCase) GetProduct(ctx context.Context, id string) (*entities.Product, error) {
	product, err := p.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductUseCase) CreateProduct(ctx context.Context, name string, price float64, stock int32) error {
	products, _ := p.repo.GetProducts(ctx)
	newProduct := &entities.Product{
		ID:        fmt.Sprint(len(products)),
		Name:      name,
		Price:     price,
		Stock:     stock,
		CreatedAt: time.Now().String(),
	}
	err := newProduct.Validate()
	if err != nil {
		return err
	}
	return p.repo.CreateProduct(ctx, newProduct)
}

func (p *ProductUseCase) UpdateProduct(ctx context.Context, id string, name *string, price *float64) error {
	updateProduct := &entities.UpdateProduct{
		Name:  name,
		Price: price,
	}
	err := updateProduct.Validate()
	if err != nil {
		return err
	}
	return p.repo.UpdateProduct(ctx, id, updateProduct)
}

func (p *ProductUseCase) DeleteProduct(ctx context.Context, id string) error {
	return p.repo.DeleteProduct(ctx, id)
}
