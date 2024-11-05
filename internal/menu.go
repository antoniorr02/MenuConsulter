package internal

import (
	"errors"
)

type DiaSemana string

const (
	Lunes     DiaSemana = "Lunes"
	Martes    DiaSemana = "Martes"
	Miercoles DiaSemana = "Miércoles"
	Jueves    DiaSemana = "Jueves"
	Viernes   DiaSemana = "Viernes"
	Sabado    DiaSemana = "Sábado"
)

type Menu struct {
	Fecha  DiaSemana
	Platos []Plato
}

func NuevoMenu(dia DiaSemana, platos []Plato) (*Menu, error) {
	if len(platos) > 5 {
		return nil, errors.New("El menú no puede tener más de 5 platos")
	}
	return &Menu{
		Fecha:  dia,
		Platos: platos,
	}, nil
}
