package repository

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/customer-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCustomerRepository struct {
	collection *mongo.Collection
}

func NewMongoCustomerRepository(client *mongo.Client, databaseName string, collectionName string) *MongoCustomerRepository {
	collection := client.Database(databaseName).Collection(collectionName)

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "customer_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, indexModel)

	if err != nil {
		log.Fatal(err)
	}

	return &MongoCustomerRepository{collection}
}

func (repository *MongoCustomerRepository) FindAllCustomers(context *gin.Context) ([]models.Customer, error) {
	cur, err := repository.collection.Find(context, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(context)

	customers := make([]models.Customer, 0)
	err = cur.All(context, &customers)

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (repository *MongoCustomerRepository) FindCustomerById(context *gin.Context, id string) (*models.Customer, error) {
	var customer models.Customer
	err := repository.collection.FindOne(context, bson.M{"customer_id": id}).Decode(&customer)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (repository *MongoCustomerRepository) CreateNewCustomer(context *gin.Context, customer *models.Customer) error {
	_, err := repository.collection.InsertOne(context, customer)

	if err != nil {
		return err
	}

	return nil
}
