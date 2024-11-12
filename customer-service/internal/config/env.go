package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var EnvProps *Env

type Env struct {
	MongoURI           string
	ServerPort         string
	CustomerDatabase   string
	CustomerCollection string
}

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	EnvProps = &Env{
		MongoURI:           os.Getenv("MONGO_URI"),
		ServerPort:         os.Getenv("SERVER_PORT"),
		CustomerDatabase:   os.Getenv("CUSTOMER_DATABASE"),
		CustomerCollection: os.Getenv("CUSTOMER_COLLECTION"),
	}
}
