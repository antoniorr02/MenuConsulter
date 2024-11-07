package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
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
	Tipo   string
	Fecha  DiaSemana
	Platos []Plato
}

func NuevoMenu(tipo string, dia DiaSemana, platos []Plato) (*Menu, error) {
	file, err := os.Open("tipos.json")
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo de tipos: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo de tipos: %w", err)
	}

	var config struct {
		Tipos []string `json:"tipos"`
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error al parsear el archivo de tipos: %w", err)
	}

	tiposPermitidos := make(map[string]bool)
	for _, t := range config.Tipos {
		tiposPermitidos[t] = true
	}

	if !tiposPermitidos[tipo] {
		return nil, errors.New("tipo de menú no permitido")
	}

	if len(platos) > 5 {
		return nil, errors.New("el menú no puede tener más de 5 platos")
	}

	return &Menu{
		Tipo:   tipo,
		Fecha:  dia,
		Platos: platos,
	}, nil
}
