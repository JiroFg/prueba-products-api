package entities

import (
	"github.com/JiroFg/prueba-products-api/internal/domain/exceptions"
)

type Product struct {
	ID        string
	Name      string
	Price     float64
	Stock     int32
	CreatedAt string
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return &exceptions.FieldError{
			FieldName: "Name",
			Message:   "Requerido",
		}
	}
	if p.Price <= 0 {
		return &exceptions.FieldError{
			FieldName: "Price",
			Message:   "No puede ser menor o igual a cero",
		}
	}
	if p.Stock <= 0 {
		return &exceptions.FieldError{
			FieldName: "Stock",
			Message:   "No puede ser menor o igual a cero",
		}
	}
	return nil
}

type UpdateProduct struct {
	Name  *string
	Price *float64
}

func (u *UpdateProduct) Validate() error {
	if u.Name != nil {
		if *u.Name == "" {
			return &exceptions.FieldError{
				FieldName: "Name",
				Message:   "No puede ser vacio",
			}
		}
	}
	if u.Price != nil {
		if *u.Price <= 0 {
			return &exceptions.FieldError{
				FieldName: "Price",
				Message:   "No puede ser menor o igual a cero",
			}
		}
	}
	return nil
}
