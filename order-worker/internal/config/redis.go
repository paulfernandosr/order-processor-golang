package config

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	clientOptions := redis.Options{
		Addr:     Props.RedisUri,
		Password: "",
		DB:       0,
	}

	redisClient := redis.NewClient(&clientOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := redisClient.Ping(ctx).Err()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successful connection to Redis")

	return redisClient
}
