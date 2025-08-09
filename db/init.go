package db

import (
	"context"
	"log"
	"pixel-battle-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(client *mongo.Client) error {
	collection := client.Database("pixels").Collection("pixels")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "index", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}

	log.Print("✅ Индекс на 'index' создан (или уже существует)")
	return nil
}

func InitBoard(client *mongo.Client) error {
	collection := client.Database("pixels").Collection("pixels")

	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	if count > 0 {
		log.Printf("✅ Коллекция 'pixels' уже содержит %d записей. Инициализация пропущена.", count)
		return nil
	}

	log.Println("🎨 Инициализация 10 000 пикселей с цветом по умолчанию (9: Белый)...")

	var board []interface{}
	for i := 1; i <= 10000; i++ {
		board = append(board, models.Pixel{
			Index: i,
			Color: 9,
		})
	}

	_, err = collection.InsertMany(context.TODO(), board)
	if err != nil {
		return err
	}
	log.Println("✅ Доска инициализирована успешно")
	return nil
}
