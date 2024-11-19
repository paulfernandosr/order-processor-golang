package repository

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type ErrorLogger interface {
	LogError(ctx context.Context, errKey string, err error) error
}

type RedisErrorLogger struct {
	client *redis.Client
}

func NewRedisErrorLogger(client *redis.Client) ErrorLogger {
	return &RedisErrorLogger{client}
}

func (errorLogger *RedisErrorLogger) LogError(ctx context.Context, errKey string, err error) error {
	log.Println(err)
	_, err = errorLogger.client.Incr(ctx, errKey).Result()
	return err
}
