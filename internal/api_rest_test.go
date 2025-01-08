package internal

import (
	"MenuConsulter/internal/config"
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

func TestGetComedores(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/comedores", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "El código de estado HTTP no es el esperado")
	config.Logger.Info("TestGetComedores ejecutado con éxito")

	assert.Contains(t, rr.Header().Get("Content-Type"), "application/json", "El tipo de contenido no es JSON")
}

func TestGetComedor(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/comedor/comedores-fuentenueva", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "El código de estado HTTP no es el esperado")
	config.Logger.Info("TestGetComedor ejecutado con éxito")

	assert.Contains(t, rr.Header().Get("Content-Type"), "application/json", "El tipo de contenido no es JSON")
}

func TestGetMenus(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/comedor/comedores-fuentenueva/menus", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "El código de estado HTTP no es el esperado")
	config.Logger.Info("TestGetMenus ejecutado con éxito")

	assert.Contains(t, rr.Header().Get("Content-Type"), "application/json", "El tipo de contenido no es JSON")
}

func TestGetMenu(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/comedor/comedores-fuentenueva/menu/2024-11-22", nil)
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
}

func TestGetPlatos(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/comedor/prueba_comedor/menu/2024-11-22/platos", nil)
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
}

func TestGetPlato(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/comedor/comedores-fuentenueva/menu/2024-11-22/plato/Manzana", nil)
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
}
