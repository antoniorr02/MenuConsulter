package internal

import (
	"MenuConsulter/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func init() {
	config.InitLogger("fichero.log")
	config.LoadConfig("config/config.yaml")
}

func cargarDocumentoTest(t *testing.T, filePath string) *html.Node {
	t.Helper()
	doc, err := cargarDocumento(filePath)
	if err != nil {
		t.Fatalf("Error al cargar el documento '%s': %v", filePath, err)
	}
	return doc
}

func encontrarTablaInline(t *testing.T, doc *html.Node) *html.Node {
	t.Helper()
	tablas := buscarNodos(doc, "table")
	for _, tabla := range tablas {
		for _, attr := range tabla.Attr {
			if attr.Key == "class" && attr.Val == "inline" {
				return tabla
			}
		}
	}
	t.Fatalf("No se encontró una tabla con la clase 'inline'")
	return nil
}

func TestExtraerFecha(t *testing.T) {
	doc := cargarDocumentoTest(t, "../data/menu.html")

	tablaInline := encontrarTablaInline(t, doc)

	fecha := extraerFecha(tablaInline)

	assert.NotEmpty(t, fecha, "No se extrajo una fecha válida")
	assert.Equal(t, "VIERNES,  22  DE  NOVIEMBRE  DE  2024", fecha, "La fecha extraída no es correcta")
}

func TestProcesarPlatos(t *testing.T) {
	doc := cargarDocumentoTest(t, "../data/menu.html")

	tablaInline := encontrarTablaInline(t, doc)

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
	config.Logger.Info("Iniciando test para ExtraerMenus", "filePath", filePath)

	menus, err := ExtraerMenus(filePath)

	// Verificar que no hay errores en la extracción
	assert.NoError(t, err, "La función ExtraerMenus devolvió un error")

	// Comprobaciones sobre el contenido
	assert.NotEmpty(t, menus, "No se extrajeron menús del archivo")
	assert.Equal(t, "Menú 1", menus[0].Tipo, "El primer menú no tiene el tipo correcto")
	assert.Contains(t, menus[0].Platos[0].Nombre, "Pastel de espinacas", "El nombre del primer plato no es correcto")
	assert.Contains(t, menus[0].Platos[0].Ingredientes, "Gluten", "No se extrajo correctamente el alérgeno del primer plato")

	config.Logger.Info("Test para ExtraerMenus completado con éxito")
}
