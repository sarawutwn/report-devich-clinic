package RedisCache

import (
	"backend-app/config"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func RegisterRedisCache() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: config.GetEnvConfig("REDIS_ADDRESS"),
	})
	return &RedisClient{client: client}
}

func (r *RedisClient) Get(key string) (string, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisClient) Set(key string, value string, exp time.Duration) error {
	err := r.client.Set(context.Background(), key, value, exp).Err()
	return err
}
