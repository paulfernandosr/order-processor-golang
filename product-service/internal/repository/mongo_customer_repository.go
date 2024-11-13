package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/paulfernandosr/order-processor-golang/product-service/internal/config"
	"github.com/paulfernandosr/order-processor-golang/product-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository interface {
	InsertOneProduct(context.Context, *model.Product) error
	FindAllProducts(context.Context) ([]model.Product, error)
	FindProductsByIds(context.Context, []string) ([]model.Product, error)
	FindOneProductById(context.Context, string) (*model.Product, error)
	UpdateOneProductById(context.Context, string, *model.Product) error
	DeleteOneProductById(context.Context, string) error
}

type MongoProductRepository struct {
	collection *mongo.Collection
}

func NewMongoProductRepository(client *mongo.Client) *MongoProductRepository {
	collection := client.Database(config.Props.ProductDatabase).Collection(config.Props.ProductCollection)

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

func (repository *MongoProductRepository) InsertOneProduct(ctx context.Context, product *model.Product) error {
	_, err := repository.collection.InsertOne(ctx, product)

	return err
}

func (repository *MongoProductRepository) FindAllProducts(ctx context.Context) ([]model.Product, error) {
	cur, err := repository.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	products := make([]model.Product, 0)
	err = cur.All(ctx, &products)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repository *MongoProductRepository) FindProductsByIds(ctx context.Context, productIds []string) ([]model.Product, error) {
	cur, err := repository.collection.Find(ctx, bson.M{"product_id": bson.M{"$in": productIds}})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	products := make([]model.Product, 0)
	err = cur.All(ctx, &products)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repository *MongoProductRepository) FindOneProductById(ctx context.Context, productId string) (*model.Product, error) {
	var product model.Product
	err := repository.collection.FindOne(ctx, bson.M{"product_id": productId}).Decode(&product)

	if err == mongo.ErrNoDocuments {
		return nil, model.NewNotFoundError(fmt.Sprintf("Product not found with identification: %s", productId))
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repository *MongoProductRepository) UpdateOneProductById(ctx context.Context, productId string, product *model.Product) error {
	product.ProductId = productId

	result, err := repository.collection.UpdateOne(ctx, bson.M{"product_id": productId}, bson.M{"$set": product})

	if result.MatchedCount == 0 {
		return model.NewNotFoundError(fmt.Sprintf("Product not found with identification: %s", productId))
	}

	return err
}

func (repository *MongoProductRepository) DeleteOneProductById(ctx context.Context, productId string) error {
	result, err := repository.collection.DeleteOne(ctx, bson.M{"product_id": productId})

	if result.DeletedCount == 0 {
		return model.NewNotFoundError(fmt.Sprintf("Product not found with identification: %s", productId))
	}

	return err
}
