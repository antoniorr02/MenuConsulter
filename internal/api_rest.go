package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

func Router() {
	if router == nil {
		router.Route("/comedores", comedoresRoutes)
	}
}

func comedoresRoutes(comedorRouter chi.Router) {
	comedorRouter.Get("/{nombre_comedor}", getcomedor)
	comedorRouter.Put("/{nombre_comedor}", putcomedor)

	comedorRouter.Route("/{nombre_comedor}/menus", menusRoutes)
}

func menusRoutes(menuRouter chi.Router) {
	menuRouter.Get("/{fecha}", getmenu)
	menuRouter.Put("/{fecha}", putmenu)

	menuRouter.Route("/{fecha}/platos", platosRoutes)
}

func platosRoutes(platoRouter chi.Router) {
	platoRouter.Get("/{nombre_plato}", getplato)
	platoRouter.Put("/{nombre_plato}", putplato)
}

func getplato(respuesta http.ResponseWriter, peticion *http.Request) {

}

func putplato(respuesta http.ResponseWriter, peticion *http.Request) {

}

func getcomedor(respuesta http.ResponseWriter, peticion *http.Request) {

}

func putcomedor(respuesta http.ResponseWriter, peticion *http.Request) {
}

func getmenu(respuesta http.ResponseWriter, peticion *http.Request) {

}

func putmenu(respuesta http.ResponseWriter, peticion *http.Request) {
}
