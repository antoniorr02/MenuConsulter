package internal

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadHTML(t *testing.T) {
	url := "https://scu.ugr.es/pages/menu/comedor"
	dataDir := "../data"
	filename := "menu.html"
	filepath := filepath.Join(dataDir, filename)

	// Crear la carpeta data si no existe
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err := os.Mkdir(dataDir, 0755)
		assert.NoError(t, err, "No se pudo crear la carpeta data")
	}

	// Realizar la petición HTTP para descargar el HTML
	resp, err := http.Get(url)
	assert.NoError(t, err, "Error al descargar el HTML")
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err, "Error al leer el contenido del HTML")

	// Guardar el contenido en el archivo
	err = os.WriteFile(filepath, body, 0644)
	assert.NoError(t, err, "Error al guardar el archivo HTML")

	// Comprobar que el archivo existe
	_, err = os.Stat(filepath)
	assert.NoError(t, err, "El archivo no se creó correctamente")

	// Comprobar que el contenido del archivo no está vacío
	content, err := os.ReadFile(filepath)
	assert.NoError(t, err, "Error al leer el archivo HTML")
	assert.NotEmpty(t, content, "El archivo HTML está vacío")

	// Eliminar el archivo después del test
	//err = os.Remove(filepath)
	//assert.NoError(t, err, "No se pudo eliminar el archivo de prueba")
}

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
