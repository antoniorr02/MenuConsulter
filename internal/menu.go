package internal

import (
	"errors"
	"time"
)

type Menu struct {
	Fecha  time.Time
	Platos []Plato
}

func NuevoMenu(fecha time.Time, platos []Plato) (*Menu, error) {
	if len(platos) > 5 {
		return nil, errors.New("El menú no puede tener más de 5 platos")
	}
	return &Menu{
		Fecha:  fecha,
		Platos: platos,
	}, nil
}
