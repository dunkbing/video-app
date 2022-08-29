package db

import (
	"context"
	"dunkbing/web-scrap/configs"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDB() {
	config := configs.GetConfig()
	dbUri := config.DbUri
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	Client = client
}

// Client instance
var Client *mongo.Client

// GetCollection getting database collections
func GetCollection(collectionName string) *mongo.Collection {
	collection := Client.Database("golangAPI").Collection(collectionName)
	return collection
}
