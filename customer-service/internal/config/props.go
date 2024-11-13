package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Props *Properties

type Properties struct {
	MongoURI           string
	ServerPort         string
	CustomerDatabase   string
	CustomerCollection string
}

func LoadEnvironment() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	Props = &Properties{
		MongoURI:           os.Getenv("MONGO_URI"),
		ServerPort:         os.Getenv("SERVER_PORT"),
		CustomerDatabase:   os.Getenv("CUSTOMER_DATABASE"),
		CustomerCollection: os.Getenv("CUSTOMER_COLLECTION"),
	}
}
