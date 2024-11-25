package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtraerMenu(t *testing.T) {
	filePath := "../data/menu.html"

	menus, err := ExtraerMenus(filePath)

	// Verificar que no hay errores en la extracción
	assert.NoError(t, err, "La función ExtraerMenus devolvió un error")

	// Comprobaciones sobre el contenido
	assert.NotEmpty(t, menus, "No se extrajeron menús del archivo")

	// Verificar algunos datos específicos de ejemplo
	assert.Equal(t, "Menú 1", menus[0].Tipo, "El primer menú no tiene el tipo correcto")
	assert.Contains(t, menus[0].Platos[0].Nombre, "Pastel de espinacas", "El nombre del primer plato no es correcto")
	assert.Contains(t, menus[0].Platos[0].Ingredientes, "Gluten", "No se extrajo correctamente el alérgeno del primer plato")
}
