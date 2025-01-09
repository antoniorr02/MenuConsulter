package internal

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func Router(router *chi.Mux, menus []Menu) *chi.Mux {
	router.Get("/menus", func(respuesta http.ResponseWriter, peticion *http.Request) {
		getMenus(respuesta, peticion, menus)
	})
	router.Get("/menu/{fecha}", func(respuesta http.ResponseWriter, peticion *http.Request) {
		getMenu(respuesta, peticion, menus)
	})
	router.Get("/menu/{fecha}/platos", func(respuesta http.ResponseWriter, peticion *http.Request) {
		getPlatos(respuesta, peticion, menus)
	})
	router.Get("/menu/{fecha}/plato/{nombre_plato}", func(respuesta http.ResponseWriter, peticion *http.Request) {
		getPlato(respuesta, peticion, menus)
	})
	return router
}

func getMenus(respuesta http.ResponseWriter, peticion *http.Request, menus []Menu) {
	respuesta.Header().Set("Content-Type", "application/json")
	json.NewEncoder(respuesta).Encode(menus)
}

func getMenu(respuesta http.ResponseWriter, peticion *http.Request, menus []Menu) {
	fecha := chi.URLParam(peticion, "fecha")

	for _, menu := range menus {
		menuFecha := string(menu.Fecha)
		if fecha == menuFecha {
			respuesta.Header().Set("Content-Type", "application/json")
			json.NewEncoder(respuesta).Encode(menu)
			return
		}
	}

	http.NotFound(respuesta, peticion)
}

func getPlatos(respuesta http.ResponseWriter, peticion *http.Request, menus []Menu) {
	fecha := chi.URLParam(peticion, "fecha")

	for _, menu := range menus {
		menuFecha := string(menu.Fecha)
		if menuFecha == fecha {
			respuesta.Header().Set("Content-Type", "application/json")
			json.NewEncoder(respuesta).Encode(menu.Platos)
			return
		}
	}

	http.NotFound(respuesta, peticion)
}

func getPlato(respuesta http.ResponseWriter, peticion *http.Request, menus []Menu) {
	fecha := chi.URLParam(peticion, "fecha")

	nombrePlato := chi.URLParam(peticion, "nombre_plato")

	for _, menu := range menus {
		menuFecha := string(menu.Fecha)
		if menuFecha == fecha {
			for _, plato := range menu.Platos {
				if strings.EqualFold(plato.Nombre, nombrePlato) {
					respuesta.Header().Set("Content-Type", "application/json")
					json.NewEncoder(respuesta).Encode(plato)
					return
				}
			}
		}
	}

	http.NotFound(respuesta, peticion)
}
