package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Abre un archivo y devuelve su contenido como un documento HTML
func cargarDocumento(filePath string) (*goquery.Document, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return nil, fmt.Errorf("error al procesar el HTML: %w", err)
	}
	return doc, nil
}

// Extrae los menús desde un documento HTML
func procesarTablas(doc *goquery.Document) []Menu {
	var menus []Menu

	doc.Find("table.inline").Each(func(i int, table *goquery.Selection) {
		fecha := extraerFecha(table)
		menus = append(menus, procesarMenuDeTabla(table, fecha)...)
	})

	return menus
}

// Extrae la fecha del menú de una tabla
func extraerFecha(table *goquery.Selection) string {
	return strings.TrimSpace(table.Find("th").First().Text())
}

// Procesa las filas de una tabla para obtener los menús
func procesarMenuDeTabla(table *goquery.Selection, fecha string) []Menu {
	var menus []Menu

	table.Find("tr").Each(func(j int, row *goquery.Selection) {
		tipo := strings.TrimSpace(row.Find("td[colspan='2']").Text())
		if tipo == "" {
			return // No es un tipo de menú
		}

		menu := Menu{
			Tipo:   tipo,
			Fecha:  DiaSemana(fecha),
			Platos: procesarPlatos(row),
		}
		menus = append(menus, menu)
	})

	return menus
}

// Procesa las filas subsecuentes de un menú para obtener los platos
func procesarPlatos(row *goquery.Selection) []Plato {
	var platos []Plato

	row.NextAll().Each(func(k int, platoRow *goquery.Selection) {
		categoria := strings.TrimSpace(platoRow.Find("td").First().Text())
		nombre := strings.TrimSpace(platoRow.Find("td").Eq(1).Text())
		alergenos := strings.TrimSpace(platoRow.Find("td").Eq(2).Text())

		if categoria == "" || nombre == "" {
			return // No es un plato válido
		}

		plato := Plato{
			Nombre:       nombre,
			Ingredientes: strings.Fields(alergenos),
		}
		platos = append(platos, plato)
	})

	return platos
}

// Función principal
func ExtraerMenus(filePath string) ([]Menu, error) {
	doc, err := cargarDocumento(filePath)
	if err != nil {
		return nil, err
	}

	return procesarTablas(doc), nil
}
