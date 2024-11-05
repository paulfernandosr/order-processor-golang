package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}
