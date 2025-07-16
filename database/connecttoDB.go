package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Replace the URI with your MongoDB connection string
	uri := "mongodb://localhost:27017/bookreviewsAPI"

	clientOpts := options.Client().ApplyURI(uri)

	var err error
	Client, err = mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}

	// Ping the database to verify connection
	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Mongo ping error:", err)
	}

	fmt.Println("✅ Connected to MongoDB")
}

func GetDB() *mongo.Database {
	return Client.Database("bookreviewsAPI")
}