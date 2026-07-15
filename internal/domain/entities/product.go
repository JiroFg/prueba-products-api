package entities

import (
	"time"
)

type Product struct {
	id        int
	name      string
	price     int // en centavos que se convertiran posteriormente a enteros dividiendo por 100
	stock     int
	createdAt time.Time
}
