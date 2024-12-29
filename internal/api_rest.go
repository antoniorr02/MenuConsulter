package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

func Router() {
	if router == nil {
		router.Route("/comedores", func(r chi.Router) {
			r.Get("/{nombre_comedor}", getcomedor)
			r.Put("/{nombre_comedor}", putcomedor)
			r.Delete("/{nombre_comedor}", deletecomedor)

			r.Route("/{nombre_comedor}/menus", func(r chi.Router) {
				r.Get("/{fecha}", getmenu)
				r.Put("/{fecha}", putmenu)
				r.Delete("/{fecha}", deletemenu)

				r.Route("/{fecha}/platos", func(r chi.Router) {
					r.Get("/{nombre_plato}", getplato)
					r.Put("/{nombre_plato}", putplato)
					r.Delete("/{nombre_plato}", deleteplato)
				})
			})
		})
	}
}

func getplato(w http.ResponseWriter, r *http.Request) {

}

func putplato(w http.ResponseWriter, r *http.Request) {

}

func deleteplato(w http.ResponseWriter, r *http.Request) {

}

func getcomedor(w http.ResponseWriter, r *http.Request) {

}

func putcomedor(w http.ResponseWriter, r *http.Request) {
}

func deletecomedor(w http.ResponseWriter, r *http.Request) {

}

func getmenu(w http.ResponseWriter, r *http.Request) {

}

func putmenu(w http.ResponseWriter, r *http.Request) {

}

func deletemenu(w http.ResponseWriter, r *http.Request) {

}
