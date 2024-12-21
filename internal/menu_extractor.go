package internal

import (
	"MenuConsulter/internal/config"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Abre un archivo y devuelve su contenido como un documento HTML
func cargarDocumento(filePath string) (*html.Node, error) {
	config.Logger.Infof("Intentando cargar el documento: %s", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		config.Logger.Infof("Error al abrir el archivo %s: %v", filePath, err)
		return nil, fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		config.Logger.Infof("Error al parsear el archivo %s: %v", filePath, err)
		return nil, fmt.Errorf("error al procesar el HTML: %w", err)
	}
	config.Logger.Infof("Documento cargado exitosamente: %s", filePath)
	return doc, nil
}

// Busca nodos en el árbol DOM por su etiqueta
func buscarNodos(n *html.Node, tagName string) []*html.Node {
	var nodes []*html.Node
	var buscar func(*html.Node)
	buscar = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tagName {
			nodes = append(nodes, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			buscar(c)
		}
	}
	buscar(n)
	return nodes
}

// Extrae el texto de un nodo y sus hijos
func extraerTexto(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(extraerTexto(c))
	}
	return sb.String()
}

// Extrae las tablas con clase "inline"
func procesarTablas(doc *html.Node) []Menu {
	config.Logger.Info("Procesando tablas en el documento")

	var menus []Menu
	tablas := buscarNodos(doc, "table")
	config.Logger.Infof("Se encontraron %d tablas", len(tablas))

	for _, tabla := range tablas {
		for _, attr := range tabla.Attr {
			if attr.Key == "class" && attr.Val == "inline" {
				config.Logger.Info("Procesando tabla con clase 'inline'")
				fecha := extraerFecha(tabla)
				menus = append(menus, procesarMenuDeTabla(tabla, fecha)...)
			}
		}
	}

	config.Logger.Infof("Se procesaron %d menús", len(menus))
	return menus
}

// Extrae la fecha del menú de una tabla
func extraerFecha(table *html.Node) string {
	thNodes := buscarNodos(table, "th")
	if len(thNodes) > 0 {
		return strings.TrimSpace(extraerTexto(thNodes[0]))
	}
	return ""
}

// Procesa las filas de una tabla para obtener los menús
func procesarMenuDeTabla(table *html.Node, fecha string) []Menu {
	config.Logger.Infof("Procesando menú para la fecha: %s", fecha)

	var menus []Menu
	trNodes := buscarNodos(table, "tr")
	config.Logger.Infof("Se encontraron %d filas en la tabla", len(trNodes))

	for _, tr := range trNodes {
		tdNodes := buscarNodos(tr, "td")
		if len(tdNodes) >= 2 {
			tipo := strings.TrimSpace(extraerTexto(tdNodes[0]))
			if tipo == "" {
				continue
			}
			menu := Menu{
				Tipo:   tipo,
				Fecha:  DiaSemana(fecha),
				Platos: procesarPlatos(tr),
			}
			menus = append(menus, menu)
			config.Logger.Infof("Se agregó un menú: %v", menu)
		}
	}
	return menus
}

// Procesa las filas subsecuentes de un menú para obtener los platos
func procesarPlatos(row *html.Node) []Plato {
	var platos []Plato
	sibling := row.NextSibling
	for sibling != nil {
		if sibling.Type == html.ElementNode && sibling.Data == "tr" {
			tdNodes := buscarNodos(sibling, "td")
			if len(tdNodes) < 3 {
				break
			}
			nombre := strings.TrimSpace(extraerTexto(tdNodes[1]))
			alergenos := strings.TrimSpace(extraerTexto(tdNodes[2]))
			if nombre == "" {
				break
			}
			plato := Plato{
				Nombre:       nombre,
				Ingredientes: strings.Fields(alergenos),
			}
			platos = append(platos, plato)
		}
		sibling = sibling.NextSibling
	}
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
