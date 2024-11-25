package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCargarDocumento(t *testing.T) {
	filePath := "../data/menu.html"

	doc, err := cargarDocumento(filePath)

	assert.NoError(t, err, "La función cargarDocumento devolvió un error")
	assert.NotNil(t, doc, "El documento cargado es nil")
}

func TestExtraerFecha(t *testing.T) {
	doc, _ := cargarDocumento("../data/menu.html")
	fecha := extraerFecha(doc.Find("table.inline").First())

	assert.NotEmpty(t, fecha, "No se extrajo una fecha válida")
	assert.Equal(t, "VIERNES,  22  DE  NOVIEMBRE  DE  2024", fecha, "La fecha extraída no es correcta")
}

func TestProcesarPlatos(t *testing.T) {
	doc, _ := cargarDocumento("../data/menu.html")
	row := doc.Find("table.inline").First().Find("tr").Eq(1)

	platos := procesarPlatos(row)

	assert.NotEmpty(t, platos, "No se procesaron platos")
	assert.Contains(t, platos[0].Nombre, "Pastel de espinacas", "El nombre del primer plato no es correcto")
	assert.Contains(t, platos[0].Ingredientes, "Gluten", "El alérgeno del primer plato no es correcto")
}

func TestExtraerMenus(t *testing.T) {
	filePath := "../data/menu.html"

	menus, err := ExtraerMenus(filePath)

	// Verificar que no hay errores en la extracción
	assert.NoError(t, err, "La función ExtraerMenus devolvió un error")

	// Comprobaciones sobre el contenido
	assert.NotEmpty(t, menus, "No se extrajeron menús del archivo")
	assert.Equal(t, "Menú 1", menus[0].Tipo, "El primer menú no tiene el tipo correcto")
	assert.Contains(t, menus[0].Platos[0].Nombre, "Pastel de espinacas", "El nombre del primer plato no es correcto")
	assert.Contains(t, menus[0].Platos[0].Ingredientes, "Gluten", "No se extrajo correctamente el alérgeno del primer plato")
}
