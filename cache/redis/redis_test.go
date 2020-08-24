package redisConnector

import (
	"github.com/go-redis/redis/v7"
	"ms.api/cache"
	"ms.api/config"
	"testing"
)

var _Cache cache.Cache

type SampleData struct {
	FirstName string
	LastName  string
}

func TestNewCache(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: config.GetSecrets().RedisURL,
		//DB:       0,
		Password: config.GetSecrets().RedisPassword,
	})
	_Cache = NewCache(client, 1000, "TEST_")
}

func TestRedisCache_Set(t *testing.T) {
	err := _Cache.Set("sample", SampleData{
		FirstName: "Justice",
		LastName:  "Nefe",
	})

	if err != nil {
		t.Log("Cannot write to cache.")
		t.Fail()
	}
}

func TestRedisCache_Get(t *testing.T) {
	s := SampleData{}
	err := _Cache.Get("sample", &s)
	if err != nil {
		t.Log("Cannot retrieve from cache.")
		t.Fail()
	}

	if s.FirstName != "Justice" {
		t.Log("Wrong data is retrieved from cache")
		t.Fail()
	}
}
