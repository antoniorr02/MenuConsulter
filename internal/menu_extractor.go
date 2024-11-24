package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

func ExtraerMenus(filePath string) ([]Menu, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer file.Close()

	// Cargar el documento HTML con goquery
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return nil, fmt.Errorf("error al procesar el HTML: %w", err)
	}

	var menus []Menu

	// Procesar las tablas de menús
	doc.Find("table.inline").Each(func(i int, table *goquery.Selection) {
		// Extraer la fecha del menú
		fecha := table.Find("th").First().Text()

		// Procesar cada menú dentro de la tabla
		table.Find("tr").Each(func(j int, row *goquery.Selection) {
			// Verificar si la fila corresponde a un tipo de menú (e.g., "Menú 1")
			tipo := row.Find("td[colspan='2']").Text()
			if tipo == "" {
				return // No es un tipo de menú, continuar
			}

			// Inicializar un nuevo menú
			menu := Menu{
				Tipo:   strings.TrimSpace(tipo),
				Fecha:  DiaSemana(strings.TrimSpace(fecha)),
				Platos: []Plato{},
			}

			// Extraer los platos del menú
			row.NextAll().Each(func(k int, platoRow *goquery.Selection) {
				categoria := platoRow.Find("td").First().Text()
				nombre := platoRow.Find("td").Eq(1).Text()
				alergenos := platoRow.Find("td").Eq(2).Text()

				if categoria == "" || nombre == "" {
					return // No es un plato válido
				}

				// Crear un nuevo plato y agregarlo al menú
				plato := Plato{
					Nombre:       strings.TrimSpace(nombre),
					Ingredientes: strings.Fields(strings.TrimSpace(alergenos)),
				}
				menu.Platos = append(menu.Platos, plato)
			})

			// Agregar el menú a la lista
			menus = append(menus, menu)
		})
	})

	return menus, nil
}
