package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Props *Properties

type Properties struct {
	MongoURI          string
	ServerPort        string
	ProductDatabase   string
	ProductCollection string
}

func LoadEnvironment() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	Props = &Properties{
		MongoURI:          os.Getenv("MONGO_URI"),
		ServerPort:        os.Getenv("SERVER_PORT"),
		ProductDatabase:   os.Getenv("PRODUCT_DATABASE"),
		ProductCollection: os.Getenv("PRODUCT_COLLECTION"),
	}
}
