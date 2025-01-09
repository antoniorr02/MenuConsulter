package internal

import (
	"MenuConsulter/internal/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.InitLogger()
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

func setupRouter() *chi.Mux {
	router := chi.NewRouter()
	filePath := "../data/menu.html"
	menus, err := ExtraerMenus(filePath)
	if err != nil {
		log.Fatalf("Error crítico al extraer menús: %v", err)
		os.Exit(1)
	}
	for i, menu := range menus {
		fechaConvertida, err := convertirFecha(string(menu.Fecha))
		if err != nil {
			log.Fatalf("Error al convertir la fecha '%s': %v", string(menu.Fecha), err)
			os.Exit(1)
		}
		menus[i].Fecha = DiaSemana(fechaConvertida)
	}

	Router(router, menus)
	return router
}

func TestGetMenus(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/menus", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "El código de estado HTTP no es el esperado")
	config.Logger.Info("TestGetMenus ejecutado con éxito")

	assert.Contains(t, rr.Header().Get("Content-Type"), "application/json", "El tipo de contenido no es JSON")

	var menus []Menu
	err = json.NewDecoder(rr.Body).Decode(&menus)
	if err != nil {
		t.Fatalf("Error al deserializar la respuesta: %v", err)
	}

	assert.NotEmpty(t, menus, "La respuesta no contiene menús")

	for _, menu := range menus {
		log.Printf("Menú del día: %s", menu.Fecha)
	}
}

func TestGetMenu(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/menu/2024-11-22", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("El código de estado HTTP fue incorrecto: %d", rr.Code)
	}
	config.Logger.Info("TestGetMenu ejecutado con éxito")

	assert.Contains(t, rr.Header().Get("Content-Type"), "application/json", "El tipo de contenido no es JSON")

	var menu Menu
	err = json.NewDecoder(rr.Body).Decode(&menu)
	if err != nil {
		t.Fatalf("Error al deserializar la respuesta: %v", err)
	}

	assert.NotEmpty(t, menu, "La respuesta no contiene ningún menú")

	log.Printf("Menú del día: %s", menu.Fecha)
}

func TestGetPlatos(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/menu/2024-11-22/platos", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("El código de estado HTTP fue incorrecto: %d", rr.Code)
	}
	config.Logger.Info("TestGetPlatos ejecutado con éxito")

	assert.Contains(t, rr.Header().Get("Content-Type"), "application/json", "El tipo de contenido no es JSON")

	var platos []Plato
	err = json.NewDecoder(rr.Body).Decode(&platos)
	if err != nil {
		t.Fatalf("Error al deserializar la respuesta: %v", err)
	}

	assert.NotEmpty(t, platos, "La respuesta no contiene platos")

	for _, plato := range platos {
		log.Printf("Menú del día: %s", plato.Nombre)
	}
}

func TestGetPlato(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/menu/2024-11-22/plato/manzana", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("El código de estado HTTP fue incorrecto: %d", rr.Code)
	}
	config.Logger.Info("TestGetPlato ejecutado con éxito")

	assert.Contains(t, rr.Header().Get("Content-Type"), "application/json", "El tipo de contenido no es JSON")

	var plato Plato
	err = json.NewDecoder(rr.Body).Decode(&plato)
	if err != nil {
		t.Fatalf("Error al deserializar la respuesta: %v", err)
	}

	assert.NotEmpty(t, plato, "La respuesta no contiene ningún plato")

	log.Printf("Menú del día: %s", plato.Nombre)
}
