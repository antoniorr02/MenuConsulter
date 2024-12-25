package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

func Router() {
	if router == nil {
		router.Get("/Plato/{Nombre}", getPlato)
		router.Put("/Plato/{Nombre}", putPlato)
		router.Delete("/Plato/{Nombre}", deletePlato)

		router.Get("/Comedor/{Nombre}", getComedor)
		router.Put("/Comedor/{Nombre}", putComedor)
		router.Delete("/Comedor/{Nombre}", deleteComedor)

		router.Get("/Menu/{Fecha}", getMenu)
		router.Put("/Menu/{Fecha}", putMenu)
		router.Delete("/Menu/{Fecha}", deleteMenu)
	}
}

func getPlato(w http.ResponseWriter, r *http.Request) {

}

func putPlato(w http.ResponseWriter, r *http.Request) {

}

func deletePlato(w http.ResponseWriter, r *http.Request) {

}

func getComedor(w http.ResponseWriter, r *http.Request) {

}

func putComedor(w http.ResponseWriter, r *http.Request) {
}

func deleteComedor(w http.ResponseWriter, r *http.Request) {

}

func getMenu(w http.ResponseWriter, r *http.Request) {

}

func putMenu(w http.ResponseWriter, r *http.Request) {

}

func deleteMenu(w http.ResponseWriter, r *http.Request) {

}
