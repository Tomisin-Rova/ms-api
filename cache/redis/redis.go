package redisConnector

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"log"
	"ms.api/cache"
	"time"
)

type redisCache struct {
	client    *redis.Client
	keyPrefix string
	ttl       time.Duration
}

func NewCache(client *redis.Client, ttl time.Duration, keyPrefix string) cache.Cache {
	if client == nil {
		log.Fatal("Redis client not connected")
	}
	return &redisCache{
		keyPrefix: keyPrefix,
		client:    client,
		ttl:       ttl,
	}
}

func (c *redisCache) Set(key string, value interface{}) error {
	input, _ := json.Marshal(value)
	_, err := c.client.Set(c.keyPrefix+key, input, c.ttl*time.Second).Result()
	return err
}

func (c *redisCache) Get(key string, output interface{}) error {
	val, err := c.client.Get(c.keyPrefix + key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), &output)
}

func (c *redisCache) Delete(key string) {
	if c.client == nil {
		return
	}
	_ = c.client.Del(c.keyPrefix + key)
}

func (c *redisCache) Update(key string, value interface{}) error {
	return c.Set(key, value)
}
