package test

import (
	"testing"

	"github.com/JiroFg/prueba-products-api/internal/repository"
	"github.com/JiroFg/prueba-products-api/internal/usecases"
)

func TestGetProductsUseCase(t *testing.T) {
	rep := repository.NewMemoryProductRepository()
	uc := usecases.NewProductUseCase(rep)
	products, err := uc.GetProducts(t.Context())
	if err != nil {
		t.Fatalf("Sucedio un error: %T", err)
	}
	if len(products) < 2 {
		t.Fatal("Menos productos de los esperados")
	}
}

func TestGetProductUseCase(t *testing.T) {
	rep := repository.NewMemoryProductRepository()
	uc := usecases.NewProductUseCase(rep)
	product, err := uc.GetProduct(t.Context(), "1")
	if err != nil {
		t.Fatalf("Sucedio un error: %T", err)
	}
	if product.Name != "Laptop" {
		t.Fatal("No se obtuvo el producto esperado")
	}
}

func TestCreateProductUseCase(t *testing.T) {
	rep := repository.NewMemoryProductRepository()
	uc := usecases.NewProductUseCase(rep)
	err := uc.CreateProduct(t.Context(), "Product1", 9.99, 10)
	if err != nil {
		t.Fatalf("Sucedio un error: %T", err)
	}
}

func TestUpdateProductUseCase(t *testing.T) {
	rep := repository.NewMemoryProductRepository()
	uc := usecases.NewProductUseCase(rep)
	name := "Producto2"
	err := uc.UpdateProduct(t.Context(), "1", &name, nil)
	if err != nil {
		t.Fatalf("Sucedio un error: %T", err)
	}
	products, err := rep.GetProducts(t.Context())
	if products[0].Name != "Producto2" {
		t.Fatal("No se actualizo el producto esperado")
	}
}

func TestDeleteProductUseCase(t *testing.T) {
	rep := repository.NewMemoryProductRepository()
	uc := usecases.NewProductUseCase(rep)
	err := uc.DeleteProduct(t.Context(), "1")
	if err != nil {
		t.Fatalf("Sucedio un error: %T", err)
	}
	products, err := rep.GetProducts(t.Context())
	if len(products) > 1 {
		t.Fatal("Los productos totales no disminuyeron")
	}
}
