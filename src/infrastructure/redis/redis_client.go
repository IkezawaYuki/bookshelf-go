package redis

import (
	"context"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/datastore"
	"github.com/go-redis/redis/v8"
	"os"
)

type client struct {
	redis *redis.Client
}

var (
	Handler datastore.RedisHandler
)

func NewRedisClient(c *redis.Client) datastore.RedisHandler {
	return &client{redis: c}
}

func GetRedisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
}

func init() {
	conn := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	Handler = &client{redis: conn}
}

func (c *client) Get(key string) (string, error) {
	return c.redis.Get(context.Background(), key).Result()
}

func (c *client) Set(key string, value interface{}) error {
	_, err := c.redis.Set(context.Background(), key, value, 0).Result()
	return err
}

func (c *client) Close() error {
	return c.redis.Close()
}

func (c *client) Delete(key string) error {
	return c.redis.Del(context.Background(), key).Err()
}
