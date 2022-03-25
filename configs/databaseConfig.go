package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectToDatabase() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database.")
	return client
}

var DB *mongo.Client = ConnectToDatabase()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Weather").Collection(collectionName)
	return collection
}
