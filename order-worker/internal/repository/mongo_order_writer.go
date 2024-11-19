package repository

import (
	"context"
	"time"

	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/config"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderWriter interface {
	Save(order *model.Order) error
}

type MongoOrderWriter struct {
	collection *mongo.Collection
}

func NewMongoOrderWriter(client *mongo.Client) OrderWriter {
	collection := client.Database(config.Props.OrderDatabase).Collection(config.Props.OrderCollection)
	return &MongoOrderWriter{collection}
}

func (orderWriter *MongoOrderWriter) Save(order *model.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := orderWriter.collection.InsertOne(ctx, order)

	return err
}
