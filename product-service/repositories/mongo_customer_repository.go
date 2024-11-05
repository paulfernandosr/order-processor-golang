package repositories

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/product-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoProductRepository struct {
	collection *mongo.Collection
}

func NewMongoProductRepository(client *mongo.Client, databaseName string, collectionName string) *MongoProductRepository {
	collection := client.Database(databaseName).Collection(collectionName)

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "product_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, indexModel)

	if err != nil {
		log.Fatal(err)
	}

	return &MongoProductRepository{collection}
}

func (repository *MongoProductRepository) FindAllProducts(context *gin.Context) ([]models.Product, error) {
	return make([]models.Product, 0), nil
}

func (repository *MongoProductRepository) FindProductById(context *gin.Context, id string) (*models.Product, error) {
	return nil, nil
}

func (repository *MongoProductRepository) CreateNewProduct(context *gin.Context, product *models.Product) error {
	_, err := repository.collection.InsertOne(context, product)

	if err != nil {
		return err
	}

	return nil
}
