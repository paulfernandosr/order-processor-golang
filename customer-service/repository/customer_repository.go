package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/customer-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository interface {
	FindAllCustomers(context *gin.Context) ([]models.Customer, error)
}

type customerMongoRepository struct {
	collection *mongo.Collection
}

func NewCustomerMongoRepository(client *mongo.Client, databaseName string, collectionName string) *customerMongoRepository {
	collection := client.Database(databaseName).Collection(collectionName)
	return &customerMongoRepository{collection}
}

func (repository *customerMongoRepository) FindAllCustomers(context *gin.Context) ([]models.Customer, error) {
	cur, err := repository.collection.Find(context, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(context)

	var customers []models.Customer
	err = cur.All(context, &customers)

	if err != nil {
		return nil, err
	}

	return customers, nil
}
