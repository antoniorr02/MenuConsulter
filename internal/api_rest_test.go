package internal

import (
	"MenuConsulter/internal/config"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.InitLogger()
}

func setupRouter() *chi.Mux {
	router := chi.NewRouter()
	Router(router)
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
