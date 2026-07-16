package test

import (
	"errors"
	"testing"
	"time"

	"github.com/JiroFg/prueba-products-api/internal/domain/entities"
	"github.com/JiroFg/prueba-products-api/internal/domain/exceptions"
	"github.com/JiroFg/prueba-products-api/internal/repository"
)

// test cuando hay dos productos que se obtuvieron correctamente
func TestGetProducts(t *testing.T) {
	repo := repository.NewMemoryProductRepository()
	products, err := repo.GetProducts(t.Context())
	if len(products) < 2 {
		t.Fatal("Hay menos de 2 productos en el repositorio")
	}
	if err != nil {
		t.Fatal("Sucedio un error")
	}
}

// test cuando obtenemos un producto en especifico por ID
func TestGetProduct(t *testing.T) {
	repo := repository.NewMemoryProductRepository()
	product, err := repo.GetProduct(t.Context(), "1")
	if product.ID != "1" {
		t.Fatal("No se obtuvo el elemento esperado")
	}
	if err != nil {
		t.Fatal(err)
	}
}

// test cuando no se encontro el producto
func TestGetProductNotFound(t *testing.T) {
	repo := repository.NewMemoryProductRepository()
	product, err := repo.GetProduct(t.Context(), "3")
	if product != nil {
		t.Fatal("Se esperaba nil")
	}
	var notFoundErr *exceptions.NotFoundError
	if !errors.As(err, &notFoundErr) {
		t.Fatalf("Se esperaba NotFoundError, se obtuvo %T", err)
	}
	if notFoundErr.ElementId != "3" {
		t.Fatalf("Se esperaba id 3, se obtuvo %s", notFoundErr.ElementId)
	}
}

// test cuando creamos un producto nuevo
func TestCreateProduct(t *testing.T) {
	repo := repository.NewMemoryProductRepository()
	newProduct := &entities.Product{
		ID:        "3",
		Name:      "Keyboard",
		Price:     19.99,
		Stock:     40,
		CreatedAt: time.Now().String(),
	}
	err := repo.CreateProduct(t.Context(), newProduct)
	if err != nil {
		t.Fatalf("Se obtuvo un error %T", err)
	}
	products, err := repo.GetProducts(t.Context())
	if len(products) < 3 {
		t.Fatal("Hay menos de 3 productos en el repositorio")
	}
}

// test cuando actualizamos un producto
func TestUpdateProduct(t *testing.T) {
	repo := repository.NewMemoryProductRepository()
	newProduct := &entities.Product{
		ID:        "3",
		Name:      "Keyboard",
		Price:     19.99,
		Stock:     40,
		CreatedAt: time.Now().String(),
	}
	err := repo.CreateProduct(t.Context(), newProduct)
	if err != nil {
		t.Fatalf("Se obtuvo un error %T", err)
	}
	products, err := repo.GetProducts(t.Context())
	if len(products) < 3 {
		t.Fatal("Hay menos de 3 productos en el repositorio")
	}
}
