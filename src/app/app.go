package app

import (
	"fmt"
	"log"
	"net/http"

	cardRoutes "src/app/routes/cardRoutes"

	"github.com/gorilla/mux"
)

func App() {
	router := mux.NewRouter()
	cardRoutes.SetUpCardRoutes(router)

	fmt.Println("server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
