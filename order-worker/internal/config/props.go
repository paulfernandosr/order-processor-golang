package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Props *Properties

type Properties struct {
	KafkaBroker            string
	OrderTopic             string
	ServerPort             string
	MongoUri               string
	OrderDatabase          string
	OrderCollection        string
	RedisUri               string
	CustomerServiceBaseUrl string
	ProductServiceBaseUrl  string
}

func LoadEnvironment() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	Props = &Properties{
		KafkaBroker:            os.Getenv("KAFKA_BROKER"),
		OrderTopic:             os.Getenv("ORDER_TOPIC"),
		ServerPort:             os.Getenv("SERVER_PORT"),
		MongoUri:               os.Getenv("MONGO_URI"),
		OrderDatabase:          os.Getenv("ORDER_DATABASE"),
		OrderCollection:        os.Getenv("ORDER_COLLECTION"),
		RedisUri:               os.Getenv("REDIS_URI"),
		CustomerServiceBaseUrl: os.Getenv("CUSTOMER_SERVICE_BASE_URL"),
		ProductServiceBaseUrl:  os.Getenv("PRODUCT_SERVICE_BASE_URL"),
	}
}
