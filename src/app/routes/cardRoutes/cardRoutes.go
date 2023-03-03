package cardroutes

import (
	cardhandlers "src/app/handlers/cardHandlers"

	"github.com/gorilla/mux"
)

func SetUpCardRoutes(router *mux.Router) {
	router.HandleFunc("/cards", cardhandlers.GetCards).Methods("GET")
	router.HandleFunc("/cards/{id}", cardhandlers.GetCardById).Methods("GET")
	router.HandleFunc("/cards", cardhandlers.CreateCard).Methods("POST")
}
