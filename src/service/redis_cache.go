package service

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type redisCache struct {
	redisClient redis.Client
}

var RedisCache = new(redisCache)

func init() {
	RedisCache.redisClient = *redis.NewClient(&redis.Options{DB: 0})
}

func (r *redisCache) Get(key string) interface{} {
	ctx := context.Background()
	result, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result
}

func (r *redisCache) GetString(key string) string {
	ctx := context.Background()
	result, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result
}

func (r *redisCache) GetInt(key string) int64 {
	ctx := context.Background()
	result, err := r.redisClient.Get(ctx, key).Int64()
	if err != nil {
		return 0
	}
	return result
}

func (r *redisCache) Incr(key string) uint64 {
	val, err := r.redisClient.Incr(context.Background(), key).Uint64()
	if err != nil {
		println(err.Error())
	}
	return val
}

func (r *redisCache) Set(key string, value interface{}) {
	ctx := context.Background()
	r.redisClient.Set(ctx, key, value, 0)
}
