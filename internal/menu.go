package internal

import (
	"errors"
	"time"
)

type Menu struct {
	Nombre string
	Fecha  time.Time
	Platos []Plato
}

func NuevoMenu(nombre string, fecha time.Time, platos []Plato) (*Menu, error) {
	if len(platos) > 5 {
		return nil, errors.New("El menú no puede tener más de 5 platos")
	}
	return &Menu{
		Nombre: nombre,
		Fecha:  fecha,
		Platos: platos,
	}, nil
}
