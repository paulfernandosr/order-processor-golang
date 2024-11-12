package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(EnvProps.MongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successful connection to MongoDB")

	return mongoClient
}
