package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexes(collection *mongo.Collection) error {
	// Create a unique compound index on BookID and UserID
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "book_id", Value: 1},
			{Key: "user_id", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Printf("Error creating unique index: %v", err)
		return err
	}

	return nil
}
