package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"MenuConsulter/internal"
)

func TestNuevoMenu_TipoInvalido(t *testing.T) {
	platos := []internal.Plato{
		{Nombre: "Plato 1", Ingredientes: []string{"Ingrediente 1"}},
	}
	_, err := internal.NuevoMenu("Desayuno", internal.Lunes, platos)
	assert.Error(t, err, "se esperaba un error por tipo inválido")
}

func TestNuevoMenu_ExcedeNumeroDePlatos(t *testing.T) {
	platos := []internal.Plato{
		{Nombre: "Plato 1", Ingredientes: []string{"Ingrediente 1"}},
		{Nombre: "Plato 2", Ingredientes: []string{"Ingrediente 2"}},
		{Nombre: "Plato 3", Ingredientes: []string{"Ingrediente 3"}},
		{Nombre: "Plato 4", Ingredientes: []string{"Ingrediente 4"}},
		{Nombre: "Plato 5", Ingredientes: []string{"Ingrediente 5"}},
		{Nombre: "Plato 6", Ingredientes: []string{"Ingrediente 6"}},
	}
	_, err := internal.NuevoMenu("Almuezo", internal.Lunes, platos)
	assert.Error(t, err, "se esperaba un error por exceder el número de platos")
}

func TestNuevoMenu_Valido(t *testing.T) {
	platos := []internal.Plato{
		{Nombre: "Plato 1", Ingredientes: []string{"Ingrediente 1"}},
		{Nombre: "Plato 2", Ingredientes: []string{"Ingrediente 2"}},
	}
	menu, err := internal.NuevoMenu("Almuerzo", internal.Lunes, platos)
	require.NoError(t, err, "no se esperaba error para un menú válido")
	assert.Equal(t, "Almuerzo", menu.Tipo)
	assert.Equal(t, internal.Lunes, menu.Fecha)
	assert.Len(t, menu.Platos, 2)
}
