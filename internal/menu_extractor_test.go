package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestCargarDocumento(t *testing.T) {
	filePath := "../data/menu.html"

	doc, err := cargarDocumento(filePath)

	assert.NoError(t, err, "La función cargarDocumento devolvió un error")
	assert.NotNil(t, doc, "El documento cargado es nil")
}

func TestExtraerFecha(t *testing.T) {
	doc, err := cargarDocumento("../data/menu.html")
	if err != nil {
		t.Fatalf("Error al cargar el documento: %v", err)
	}

	// Busca la primera tabla con clase "inline"
	tablas := buscarNodos(doc, "table")
	var tablaInline *html.Node
	for _, tabla := range tablas {
		for _, attr := range tabla.Attr {
			if attr.Key == "class" && attr.Val == "inline" {
				tablaInline = tabla
				break
			}
		}
		if tablaInline != nil {
			break
		}
	}

	if tablaInline == nil {
		t.Fatalf("No se encontró una tabla con la clase 'inline'")
	}

	fecha := extraerFecha(tablaInline)

	assert.NotEmpty(t, fecha, "No se extrajo una fecha válida")
	assert.Equal(t, "VIERNES,  22  DE  NOVIEMBRE  DE  2024", fecha, "La fecha extraída no es correcta")
}

func TestProcesarPlatos(t *testing.T) {
	doc, err := cargarDocumento("../data/menu.html")
	if err != nil {
		t.Fatalf("Error al cargar el documento: %v", err)
	}

	// Busca la primera tabla con clase "inline"
	tablas := buscarNodos(doc, "table")
	var tablaInline *html.Node
	for _, tabla := range tablas {
		for _, attr := range tabla.Attr {
			if attr.Key == "class" && attr.Val == "inline" {
				tablaInline = tabla
				break
			}
		}
		if tablaInline != nil {
			break
		}
	}

	if tablaInline == nil {
		t.Fatalf("No se encontró una tabla con la clase 'inline'")
	}

	// Encuentra la segunda fila de la tabla
	filas := buscarNodos(tablaInline, "tr")
	if len(filas) < 2 {
		t.Fatalf("No se encontraron suficientes filas en la tabla")
	}

	row := filas[1]
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
