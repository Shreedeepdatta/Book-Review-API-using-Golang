package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateAllIndexes() {
	db := GetDB()

	// User indexes
	userCollection := db.Collection("users")
	createUserIndexes(userCollection)

	// Author indexes
	authorCollection := db.Collection("authors")
	createAuthorIndexes(authorCollection)

	// Book indexes
	bookCollection := db.Collection("books")
	createBookIndexes(bookCollection)

	// Review indexes
	reviewCollection := db.Collection("reviews")
	createReviewIndexes(reviewCollection)
}

func createUserIndexes(collection *mongo.Collection) {
	// Unique username index
	usernameIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	// Unique email index
	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{usernameIndex, emailIndex})
	if err != nil {
		log.Printf("Error creating user indexes: %v", err)
	}
}

func createAuthorIndexes(collection *mongo.Collection) {
	// User ID index
	userIDIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "user_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), userIDIndex)
	if err != nil {
		log.Printf("Error creating author indexes: %v", err)
	}
}

func createBookIndexes(collection *mongo.Collection) {
	// Author ID index
	authorIDIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "author_id", Value: 1}},
	}

	// Genre index
	genreIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "genre", Value: 1}},
	}

	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{authorIDIndex, genreIndex})
	if err != nil {
		log.Printf("Error creating book indexes: %v", err)
	}
}

func createReviewIndexes(collection *mongo.Collection) error {
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
