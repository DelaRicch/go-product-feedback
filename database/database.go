package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDb() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	mongoURI := fmt.Sprintf("mongodb+srv://delaricch:%s@product-feedback.l8phg4q.mongodb.net/?retryWrites=true&w=majority", os.Getenv("MONGO_PASSWORD"))
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	return nil
}

func DisconnectDb() error {
	err := client.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
    // Check if database exists, create if not
    if err := client.Database("product-feedback-db").CreateCollection(context.Background(), collectionName); err != nil && !strings.Contains(err.Error(), "NamespaceExists") {
        log.Fatal(err)
    }

    return client.Database("product-feedback-db").Collection(collectionName)
}