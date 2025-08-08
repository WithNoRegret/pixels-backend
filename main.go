package main

import (
	"log"
	"net/http"

	_ "pixel-battle-backend/docs"
	"pixel-battle-backend/handlers"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Pixel Battle API
// @version 1.0
// @description API for Pixel Battle game
// @host backend.battling-pixels.ru
// @BasePath /
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/palette", handlers.PaletteHandler).Methods("GET")
	router.HandleFunc("/palette/", handlers.PaletteHandler).Methods("GET")

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler()).Methods("GET")
	
	corsHandler := ghandlers.CORS(
		ghandlers.AllowedOrigins([]string{"*"}),
		ghandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		ghandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(router)))
}
