package entities

import "errors"

type Product struct {
	ID        string
	Name      string
	Price     float64
	Stock     int32
	CreatedAt string
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("Name requerido")
	}
	if p.Price <= 0 {
		return errors.New("Price no puede ser menor o igual a cero")
	}
	if p.Stock < 0 {
		return errors.New("Stock no puede ser menor o igual a cero")
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
			return errors.New("Name no puede ser vacio")
		}
	}
	if u.Price != nil {
		if *u.Price <= 0 {
			return errors.New("Price no puede ser menor o igual a cero")
		}
	}
	return nil
}
