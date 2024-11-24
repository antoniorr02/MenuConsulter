package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadHTML(url string, destDir string, filename string) error {
	// Crear la carpeta de destino si no existe
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err := os.Mkdir(destDir, 0755)
		if err != nil {
			return fmt.Errorf("no se pudo crear la carpeta %s: %v", destDir, err)
		}
	}

	// Realizar la petición HTTP
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error al descargar el contenido desde %s: %v", url, err)
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error al leer el contenido descargado: %v", err)
	}

	// Crear la ruta completa del archivo
	filepath := filepath.Join(destDir, filename)

	// Guardar el contenido en el archivo
	err = os.WriteFile(filepath, body, 0644)
	if err != nil {
		return fmt.Errorf("error al guardar el archivo %s: %v", filepath, err)
	}

	return nil
}
