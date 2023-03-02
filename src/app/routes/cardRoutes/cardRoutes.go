package cardroutes

import (
	cardhandlers "src/app/handlers/cardHandlers"

	"github.com/gorilla/mux"
)

func SetUpRoutes(router *mux.Router) {
	router.HandleFunc("/cards", cardhandlers.GetCards).Methods("GET")
}
