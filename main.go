package main

import (
	"log"
	"net/http"

	"pixel-battle-backend/db"
	"pixel-battle-backend/handlers"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	client := db.ConnectToMongo()
	err := db.InitDB(client)
	if err != nil {
		log.Fatal("❌ Ошибка при инициализации БД:", err)
	}
	err = db.InitBoard(client)
	if err != nil {
		log.Fatal("❌ Ошибка при инициализации доски:", err)
	}

	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/palette", handlers.PaletteHandler).Methods("GET")

	boardHandler := handlers.BoardHandler(client)
	router.HandleFunc("/board", boardHandler).Methods("GET")

	corsHandler := ghandlers.CORS(
		ghandlers.AllowedOrigins([]string{"*"}),
		ghandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		ghandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(router)))
}
