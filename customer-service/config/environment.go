package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var EnvironmentProps *Environment

type Environment struct {
	MongoURI   string
	ServerPort string
}

func LoadEnviroment() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	EnvironmentProps = &Environment{
		MongoURI:   os.Getenv("MONGO_URI"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
