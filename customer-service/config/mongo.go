package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(os.Getenv(EnvironmentProps.MongoURI))

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	MongoClient, err := mongo.Connect(context, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = MongoClient.Ping(context, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successful connection to MongoDB")

	return MongoClient
}
