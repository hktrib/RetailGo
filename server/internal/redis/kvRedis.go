package kv

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/hktrib/RetailGo/internal/ent"
)

type Cache struct {
	Client        *redis.Client
	defaultExpiry time.Duration
	ctx           context.Context
	opts          *redis.Options
}

func NewCache(ctx context.Context, opts *redis.Options, defaultExpiry time.Duration) *Cache {
	cache := &Cache{}

	cache.Client = cache.getClient(opts)
	cache.defaultExpiry = defaultExpiry
	cache.ctx = ctx
	cache.opts = opts
	return cache
}

func (c *Cache) getClient(opts *redis.Options) *redis.Client {
	return redis.NewClient(opts)
}

func (c *Cache) Set(key string, value *ent.User) {
	client := c.Client
	// serialize value object to JSON

	fmt.Printf("Set Key: %s\n", key)
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(c.ctx, key, json, 0)
}

func (c *Cache) SetX(key string, value *ent.User, expiresAt time.Duration) {
	client := c.Client

	// serialize value object to JSON
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(c.ctx, key, json, expiresAt)
}

func (c *Cache) Get(key string) *ent.User {
	client := c.Client

	fmt.Printf("Get Key: %s\n", key)
	ctx := context.Background()
	val, err := client.Get(ctx, key).Result()
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
