package kvRedis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/hktrib/RetailGo/ent"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
	ctx     context.Context
}

func NewRedisCache(host string, db int, exp time.Duration) *redisCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
		ctx:     context.Background(),
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, entity *ent.User) {
	client := cache.getClient()

	// serialize value object to JSON
	json, err := json.Marshal(entity)
	if err != nil {
		panic(err)
	}

	client.Set(cache.ctx, key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *ent.User {
	client := cache.getClient()

	val, err := client.Get(cache.ctx, key).Result()
	if err != nil {
		return nil
	}

	entity := ent.User{}
	err = json.Unmarshal([]byte(val), &entity)
	if err != nil {
		panic(err)
	}

	return &entity
}
