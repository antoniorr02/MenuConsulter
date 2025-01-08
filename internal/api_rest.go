package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

func Router(router *chi.Mux) *chi.Mux {
	router.Get("/comedores", getComedores)
	router.Get("/comedor/{nombre_comedor}", getComedor)
	router.Get("/comedor/{nombre_comedor}/menus", getMenus)
	router.Get("/comedor/{nombre_comedor}/menu/{fecha}", getMenu)
	router.Get("/comedor/{nombre_comedor}/menu/{fecha}/platos", getPlatos)
	router.Get("/comedor/{nombre_comedor}/menu/{fecha}/plato/{nombre_plato}", getPlato)
	return router
}

func getComedores(respuesta http.ResponseWriter, peticion *http.Request) {
	filePath := "../data/menu.html"
	menus, err := ExtraerMenus(filePath)
	if err != nil {
		log.Fatalf("Error crítico al extraer menús: %v", err)
		os.Exit(1)
	}

	var comedores []string
	for _, menu := range menus {
		if !contains(comedores, menu.Tipo) {
			comedores = append(comedores, menu.Tipo)
		}
	}

	respuesta.Header().Set("Content-Type", "application/json")
	json.NewEncoder(respuesta).Encode(comedores)
}

func getComedor(respuesta http.ResponseWriter, peticion *http.Request) {
	nombreComedor := chi.URLParam(peticion, "nombre_comedor")
	filePath := "../data/menu.html"
	menus, err := ExtraerMenus(filePath)
	if err != nil {
		log.Fatalf("Error crítico al extraer menús: %v", err)
		os.Exit(1)
	}

	var comedorMenus []Menu
	for _, menu := range menus {
		if menu.Tipo == nombreComedor {
			comedorMenus = append(comedorMenus, menu)
		}
	}

	respuesta.Header().Set("Content-Type", "application/json")
	json.NewEncoder(respuesta).Encode(comedorMenus)
}

func getMenus(respuesta http.ResponseWriter, peticion *http.Request) {
	nombreComedor := chi.URLParam(peticion, "nombre_comedor")
	filePath := "../data/menu.html"
	menus, err := ExtraerMenus(filePath)
	if err != nil {
		log.Fatalf("Error crítico al extraer menús: %v", err)
		os.Exit(1)
	}

	var filteredMenus []Menu
	for _, menu := range menus {
		if menu.Tipo == nombreComedor {
			filteredMenus = append(filteredMenus, menu)
		}
	}

	respuesta.Header().Set("Content-Type", "application/json")
	json.NewEncoder(respuesta).Encode(filteredMenus)
}

func convertirFecha(fecha string) (string, error) {
	fecha = strings.TrimSpace(fecha)
	fecha = strings.ReplaceAll(fecha, "  ", " ")

	if _, err := time.Parse("2006-01-02", fecha); err == nil {
		return fecha, nil
	}

	parts := strings.Split(fecha, ",")
	if len(parts) < 2 {
		return "", fmt.Errorf("fecha en formato incorrecto: %s", fecha)
	}
	fecha = strings.TrimSpace(parts[1])

	partes := strings.Fields(fecha)
	if len(partes) < 5 {
		return "", fmt.Errorf("fecha incompleta: %s", fecha)
	}

	dia := partes[0]
	mes := partes[2]
	ano := partes[4]

	meses := map[string]string{
		"enero":      "01",
		"febrero":    "02",
		"marzo":      "03",
		"abril":      "04",
		"mayo":       "05",
		"junio":      "06",
		"julio":      "07",
		"agosto":     "08",
		"septiembre": "09",
		"octubre":    "10",
		"noviembre":  "11",
		"diciembre":  "12",
	}

	mes = strings.ToLower(mes)
	mesNumero, ok := meses[mes]
	if !ok {
		return "", fmt.Errorf("mes no válido: %s", mes)
	}

	fechaFormateada := fmt.Sprintf("%s-%s-%s", ano, mesNumero, dia)
	return fechaFormateada, nil
}

func getMenu(respuesta http.ResponseWriter, peticion *http.Request) {
	fecha := chi.URLParam(peticion, "fecha")
	filePath := "../data/menu.html"
	menus, err := ExtraerMenus(filePath)
	if err != nil {
		log.Fatalf("Error crítico al extraer menús: %v", err)
		os.Exit(1)
	}

	fechaFormateada, err := convertirFecha(fecha)
	if err != nil {
		http.Error(respuesta, "Fecha inválida", http.StatusBadRequest)
		log.Printf("Error en la conversión de fecha: %s", err)
		return
	}

	for _, menu := range menus {
		menuFecha := string(menu.Fecha)
		menuFechaFormateada, err := convertirFecha(menuFecha)
		if err != nil {
			http.Error(respuesta, "Fecha inválida", http.StatusBadRequest)
			log.Printf("Error en la conversión de fecha: %s", err)
			return
		}

		if menuFechaFormateada == fechaFormateada {
			respuesta.Header().Set("Content-Type", "application/json")
			json.NewEncoder(respuesta).Encode(menu)
			return
		}
	}

	http.NotFound(respuesta, peticion)
}

func getPlatos(respuesta http.ResponseWriter, peticion *http.Request) {
	fecha := chi.URLParam(peticion, "fecha")
	filePath := "../data/menu.html"
	menus, err := ExtraerMenus(filePath)
	if err != nil {
		log.Fatalf("Error crítico al extraer menús: %v", err)
		os.Exit(1)
	}

	fechaFormateada, err := convertirFecha(fecha)
	if err != nil {
		http.Error(respuesta, "Fecha inválida", http.StatusBadRequest)
		return
	}

	for _, menu := range menus {
		menuFecha := string(menu.Fecha)
		menuFechaFormateada, err := convertirFecha(menuFecha)
		if err != nil {
			http.Error(respuesta, "Fecha inválida", http.StatusBadRequest)
			log.Printf("Error en la conversión de fecha: %s", err)
			return
		}
		if menuFechaFormateada == fechaFormateada {
			respuesta.Header().Set("Content-Type", "application/json")
			json.NewEncoder(respuesta).Encode(menu.Platos)
			return
		}
	}

	http.NotFound(respuesta, peticion)
}

func getPlato(respuesta http.ResponseWriter, peticion *http.Request) {
	fecha := chi.URLParam(peticion, "fecha")
	nombrePlato := chi.URLParam(peticion, "nombre_plato")
	filePath := "../data/menu.html"
	menus, err := ExtraerMenus(filePath)
	if err != nil {
		log.Fatalf("Error crítico al extraer menús: %v", err)
		os.Exit(1)
	}

	fechaFormateada, err := convertirFecha(fecha)
	if err != nil {
		http.Error(respuesta, "Fecha inválida", http.StatusBadRequest)
		return
	}

	for _, menu := range menus {
		menuFecha := string(menu.Fecha)
		menuFechaFormateada, err := convertirFecha(menuFecha)
		if err != nil {
			http.Error(respuesta, "Fecha inválida", http.StatusBadRequest)
			log.Printf("Error en la conversión de fecha: %s", err)
			return
		}
		if menuFechaFormateada == fechaFormateada {
			for _, plato := range menu.Platos {
				if plato.Nombre == nombrePlato {
					respuesta.Header().Set("Content-Type", "application/json")
					json.NewEncoder(respuesta).Encode(plato)
					return
				}
			}
		}
	}

	http.NotFound(respuesta, peticion)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
