package redis_client

import (
	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	redis *redis.Client
}

func (c *redisClient) Get() {

}
