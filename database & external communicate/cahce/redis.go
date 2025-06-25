package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisCahce struct {
	client *redis.Client
}

func NewRedisCahce() *RedisCahce {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if client == nil {
		fmt.Println("Redis client is nil")
	}

	return &RedisCahce{
		client: client,
	}
}

var ctx = context.Background()

func (r RedisCahce) Get() []byte {
	strCmd := r.client.Get(ctx, "cities")
	cacheBytes, _ := strCmd.Bytes()
	return cacheBytes
}

func (r RedisCahce) Put(value []byte) {
	r.client.Set(ctx, "cities", value, 0)
}
