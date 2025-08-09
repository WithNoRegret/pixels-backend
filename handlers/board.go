package handlers

import (
	"context"
	"log"

	"pixel-battle-backend/api"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBoardData(client *mongo.Client) ([]api.Pixel, error) {
	collection := client.Database("pixels").Collection("pixels")

	cursor, err := collection.Find(context.Background(), bson.M{},
		options.Find().SetSort(bson.D{{Key: "index", Value: 1}}))
	if err != nil {
		log.Println("❌ Ошибка при получении доски:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var board []api.Pixel
	err = cursor.All(context.Background(), &board)
	if err != nil {
		log.Println("❌ Ошибка при парсинге данных:", err)
		return nil, err
	}

	return board, nil
}
