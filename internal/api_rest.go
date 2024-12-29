package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

func Router() {
	if router == nil {
		router.Get("/platos/{nombre}", getplato)
		router.Put("/platos/{nombre}", putplato)
		router.Delete("/platos/{nombre}", deleteplato)

		router.Get("/comedores/{nombre}", getcomedor)
		router.Put("/comedores/{nombre}", putcomedor)
		router.Delete("/comedores/{nombre}", deletecomedor)

		router.Get("/menus/{fecha}", getmenu)
		router.Put("/menus/{fecha}", putmenu)
		router.Delete("/menus/{fecha}", deletemenu)
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
