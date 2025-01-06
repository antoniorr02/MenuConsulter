package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
	Router()
}

func Router() {
	router.Get("/comedores", getComedores)
	router.Get("/comedor/{nombre_comedor}", getComedor)
	router.Get("/comedores/{nombre_comedor}/menus", getMenus)
	router.Get("/comedores/{nombre_comedor}/menu/{fecha}", getMenu)
	router.Get("/comedores/{nombre_comedor}/menus/{fecha}/plato/{nombre_plato}", getPlato)
}

func getPlato(respuesta http.ResponseWriter, peticion *http.Request) {

}

func getComedores(respuesta http.ResponseWriter, peticion *http.Request) {

}

func getComedor(respuesta http.ResponseWriter, peticion *http.Request) {

}

func getMenu(respuesta http.ResponseWriter, peticion *http.Request) {

}

func getMenus(respuesta http.ResponseWriter, peticion *http.Request) {

}
