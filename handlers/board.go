package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"pixel-battle-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BoardHandler(client *mongo.Client) http.HandlerFunc {
	collection := client.Database("pixels").Collection("pixels")

	return func(w http.ResponseWriter, r *http.Request) {
		cursor, err := collection.Find(r.Context(), bson.M{}, options.Find().SetSort(bson.D{{Key: "index", Value: 1}}))
		if err != nil {
			log.Println("❌ Ошибка при получении доски:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer cursor.Close(r.Context())

		var board []models.Pixel
		err = cursor.All(r.Context(), &board)
		if err != nil {
			log.Println("❌ Ошибка при парсинге данных:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(board)
		if err != nil {
			log.Println("❌ Ошибка при отправке json:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
