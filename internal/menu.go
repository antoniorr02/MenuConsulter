package internal

import (
	"time"
)

type Menu struct {
	Nombre string
	Fecha time.Time
	Platos []Plato
}