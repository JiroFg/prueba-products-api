package test

import (
	"testing"
	"time"

	"github.com/JiroFg/prueba-products-api/internal/domain/entities"
)

// test cuando creamos un producto nuevo correctamente
func TestNewProduct(t *testing.T) {
	newProduct := &entities.Product{
		Name:      "Producto1",
		Price:     9.99,
		Stock:     10,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	err := newProduct.Validate()
	if err != nil {
		t.Fatal("Error al crear el producto")
	}
}

func TestNewProductInvalidName(t *testing.T) {
	newProduct := &entities.Product{}
	err := newProduct.Validate()
	if err == nil {
		t.Fatal("El producto se creo correctamente")
	}
	if err.Error() != "Campo: \"Name\" con error: \"Requerido\"" {
		t.Fatal("Error diferente al esperado")
	}
}

func TestNewProductInvalidPrice(t *testing.T) {
	newProduct := &entities.Product{
		Name: "Product1",
	}
	err := newProduct.Validate()
	if err == nil {
		t.Fatal("El producto se creo correctamente")
	}
	if err.Error() != "Campo: \"Price\" con error: \"No puede ser menor o igual a cero\"" {
		t.Fatal("Error diferente al esperado")
	}
}

func TestNewProductInvalidStock(t *testing.T) {
	newProduct := &entities.Product{
		Name:  "Product1",
		Price: 9.99,
		Stock: 0,
	}
	err := newProduct.Validate()
	if err == nil {
		t.Fatal("El producto se creo correctamente")
	}
	if err.Error() != "Campo: \"Stock\" con error: \"No puede ser menor o igual a cero\"" {
		t.Fatal("Error diferente al esperado")
	}
}
