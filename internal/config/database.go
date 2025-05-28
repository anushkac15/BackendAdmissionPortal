package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI environment variable is not set")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)
	
	// Add additional options for MongoDB Atlas
	clientOptions.SetServerSelectionTimeout(5 * time.Second)
	clientOptions.SetConnectTimeout(10 * time.Second)
	clientOptions.SetSocketTimeout(10 * time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	// Set the database
	DB = client.Database("admission_portal")
	log.Println("Connected to MongoDB Atlas!")
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
} 