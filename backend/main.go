package main

import (
	"log"
	"net/http"

	_ "pixel-battle-backend/docs"
	"pixel-battle-backend/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
					w.WriteHeader(http.StatusOK)
					return
			}

			next.ServeHTTP(w, r)
	}
}


// @title Pixel Battle API
// @version 1.0
// @description API for Pixel Battle game
// @host backend.battling-pixels.ru
// @BasePath /
func main() {
	http.HandleFunc("/", CorsMiddleware(handlers.HomeHandler))
	http.HandleFunc("/palette/", CorsMiddleware(handlers.PaletteHandler))

	http.HandleFunc("/swagger/", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
