package config

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

func init() {
	router = chi.NewRouter()
	Router()
}

func Router() {
	router.Get("/comedores/{nombre_comedor}", getComedor)
	router.Get("/comedores/{nombre_comedor}/menus/{fecha}", getMenu)
	router.Get("/comedores/{nombre_comedor}/menus/{fecha}/platos/{nombre_plato}", getPlato)
}

func getPlato(respuesta http.ResponseWriter, peticion *http.Request) {

}

func getComedor(respuesta http.ResponseWriter, peticion *http.Request) {

}

func getMenu(respuesta http.ResponseWriter, peticion *http.Request) {

}
