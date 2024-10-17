package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache struct {
	rdb            *redis.Client
	expirationTime time.Duration
}

func New(rdb *redis.Client, expirationTime time.Duration) *RedisCache {
	return &RedisCache{
		rdb:            rdb,
		expirationTime: expirationTime,
	}
}

func (cache *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return cache.rdb.Get(ctx, key).Result()
}

func (cache *RedisCache) Set(ctx context.Context, key, value string) error {
	return cache.rdb.Set(ctx, key, value, cache.expirationTime).Err()
}
