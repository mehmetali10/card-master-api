package app

import (
	"fmt"
	"log"
	"net/http"

	cardRoutes "src/app/routes/cardRoutes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func App() {
	router := mux.NewRouter()
	cardRoutes.SetUpCardRoutes(router)

	// Cors ayarlarını yapmak için bir cors seçenekleri oluşturuyoruz
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Sadece http://localhost:3000 kaynağından gelen isteklere izin ver
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"}, // Tüm başlıklara izin ver
	})

	// Cors özelliklerini uygulamak için http.Handler'ı sarmalıyoruz
	handler := corsOptions.Handler(router)

	fmt.Println("server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
