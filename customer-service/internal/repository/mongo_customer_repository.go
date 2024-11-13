package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/config"
	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CustomerRepository interface {
	InsertOneCustomer(context.Context, *model.Customer) error
	FindAllCustomers(context.Context) ([]model.Customer, error)
	FindCustomersByIds(context.Context, []string) ([]model.Customer, error)
	FindOneCustomerById(context.Context, string) (*model.Customer, error)
	UpdateOneCustomerById(context.Context, string, *model.Customer) error
	DeleteOneCustomerById(context.Context, string) error
}

type MongoCustomerRepository struct {
	collection *mongo.Collection
}

func NewMongoCustomerRepository(client *mongo.Client) *MongoCustomerRepository {
	collection := client.Database(config.Props.CustomerDatabase).Collection(config.Props.CustomerCollection)

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

func (repository *MongoCustomerRepository) InsertOneCustomer(ctx context.Context, customer *model.Customer) error {
	_, err := repository.collection.InsertOne(ctx, customer)

	return err
}

func (repository *MongoCustomerRepository) FindAllCustomers(ctx context.Context) ([]model.Customer, error) {
	cur, err := repository.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	customers := make([]model.Customer, 0)
	err = cur.All(ctx, &customers)

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (repository *MongoCustomerRepository) FindCustomersByIds(ctx context.Context, customerIds []string) ([]model.Customer, error) {
	cur, err := repository.collection.Find(ctx, bson.M{"customer_id": bson.M{"$in": customerIds}})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	customers := make([]model.Customer, 0)
	err = cur.All(ctx, &customers)

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (repository *MongoCustomerRepository) FindOneCustomerById(ctx context.Context, customerId string) (*model.Customer, error) {
	var customer model.Customer
	err := repository.collection.FindOne(ctx, bson.M{"customer_id": customerId}).Decode(&customer)

	if err == mongo.ErrNoDocuments {
		return nil, model.NewNotFoundError(fmt.Sprintf("Customer not found with identification: %s", customerId))
	}

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (repository *MongoCustomerRepository) UpdateOneCustomerById(ctx context.Context, customerId string, customer *model.Customer) error {
	customer.CustomerId = customerId

	result, err := repository.collection.UpdateOne(ctx, bson.M{"customer_id": customerId}, bson.M{"$set": customer})

	if result.MatchedCount == 0 {
		return model.NewNotFoundError(fmt.Sprintf("Customer not found with identification: %s", customerId))
	}

	return err
}

func (repository *MongoCustomerRepository) DeleteOneCustomerById(ctx context.Context, customerId string) error {
	result, err := repository.collection.DeleteOne(ctx, bson.M{"customer_id": customerId})

	if result.DeletedCount == 0 {
		return model.NewNotFoundError(fmt.Sprintf("Customer not found with identification: %s", customerId))
	}

	return err
}
