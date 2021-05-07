package redis_client

import (
	"context"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/datastore"
	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	redis *redis.Conn
}

func NewRedisClient(c *redis.Conn) datastore.RedisHandler {
	return &redisClient{redis: c}
}

func (c *redisClient) Get(key string) (string, error) {
	return c.redis.Get(context.Background(), key).Result()
}

func (c *redisClient) Set(key string, value interface{}) (string, error) {
	return c.redis.Set(context.Background(), key, value, 0).Result()
}

func (c *redisClient) Close() error {
	return c.redis.Close()
}
