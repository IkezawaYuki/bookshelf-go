package infrastructure

import (
	"context"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/datastore"
	"github.com/go-redis/redis/v8"
	"os"
)

type redisClient struct {
	redis *redis.Client
}

var (
	RedisHandler datastore.RedisHandler
)

func NewRedisClient(c *redis.Client) datastore.RedisHandler {
	return &redisClient{redis: c}
}

func GetRedisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0, // use default DB
	})
}

func init() {
	conn := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0, // use default DB
	})
	RedisHandler = &redisClient{redis: conn}
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

func (c *redisClient) Delete(key string) error {
	return c.redis.Del(context.Background(), key).Err()
}
