package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"pixel-battle-backend/constants"

	"go.mongodb.org/mongo-driver/mongo"
)

type APIImpl struct {
	Client *mongo.Client
}

func (a *APIImpl) GetBoard(w http.ResponseWriter, r *http.Request) {
	board, err := GetBoardData(a.Client)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(board); err != nil {
		log.Println("❌ Ошибка при отправке json:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (a *APIImpl) GetPalette(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(constants.ColorsPalette); err != nil {
		log.Println("❌ Ошибка при отправке json:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
