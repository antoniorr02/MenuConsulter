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
	Domingo   DiaSemana = "Domingo"
)

type Mes string

const (
	Enero      Mes = "Enero"
	Febrero    Mes = "Febrero"
	Marzo      Mes = "Marzo"
	Abril      Mes = "Abril"
	Mayo       Mes = "Mayo"
	Junio      Mes = "Junio"
	Julio      Mes = "Julio"
	Agosto     Mes = "Agosto"
	Septiembre Mes = "Septiembre"
	Octubre    Mes = "Octubre"
	Noviembre  Mes = "Noviembre"
	Diciembre  Mes = "Diciembre"
)

type FechaCompleta struct {
	DiaSemana DiaSemana
	Dia       int
	Mes       Mes
	Anio      int
}

func CrearFechaCompleta(diaSemana DiaSemana, dia int, mes Mes, anio int) (FechaCompleta, error) {
	if dia < 1 || dia > 31 {
		return FechaCompleta{}, errors.New("día inválido")
	}
	if anio < 1 {
		return FechaCompleta{}, errors.New("año inválido")
	}
	if !mesValido(mes, dia, anio) {
		return FechaCompleta{}, errors.New("mes o día inválido para el mes y año dado")
	}

	return FechaCompleta{
		DiaSemana: diaSemana,
		Dia:       dia,
		Mes:       mes,
		Anio:      anio,
	}, nil
}

func esBisiesto(anio int) bool {
	return (anio%4 == 0 && anio%100 != 0) || (anio%400 == 0)
}

func mesValido(mes Mes, dia int, anio int) bool {
	switch mes {
	case Abril, Junio, Septiembre, Noviembre:
		return dia <= 30
	case Febrero:
		if esBisiesto(anio) {
			return dia <= 29
		}
		return dia <= 28
	default:
		return dia <= 31
	}
}

type Menu struct {
	Fecha  FechaCompleta
	Platos []Plato
}

func NuevoMenu(fecha FechaCompleta, platos []Plato) (*Menu, error) {
	if len(platos) > 5 {
		return nil, errors.New("El menú no puede tener más de 5 platos")
	}
	return &Menu{
		Fecha:  fecha,
		Platos: platos,
	}, nil
}
