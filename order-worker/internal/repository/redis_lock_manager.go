package repository

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type LockManager interface {
	AcquireLock(ctx context.Context, lockKey string, clientId string, ttl time.Duration) (bool, error)
	ReleaseLock(ctx context.Context, lockKey string, clientId string) (bool, error)
}

type RedisLockManager struct {
	client *redis.Client
}

func NewRedisLockManager(client *redis.Client) LockManager {
	return &RedisLockManager{client}
}

func (lockManager *RedisLockManager) AcquireLock(ctx context.Context, lockKey string, clientId string, ttl time.Duration) (bool, error) {
	return lockManager.client.SetNX(ctx, lockKey, clientId, ttl).Result()
}

func (lockManager *RedisLockManager) ReleaseLock(ctx context.Context, lockKey string, clientId string) (bool, error) {
	script := redis.NewScript(`
		if redis.call("GET", KEYS[1]) ~= ARGV[1] then
			return 0
		end

		return redis.call("DEL", KEYS[1])
	`)

	result, err := script.Run(ctx, lockManager.client, []string{lockKey}, clientId).Result()

	if err != nil {
		return false, err
	}

	scriptResult, ok := result.(int64)

	if !ok {
		return false, errors.New("invalid script result")
	}

	return scriptResult == 1, nil
}
