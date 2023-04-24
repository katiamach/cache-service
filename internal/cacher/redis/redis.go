package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

// Cacher wraps redis client.
type Cacher struct {
	redis *redis.Client
}

// New creates new cacher for the service.
func New() (*Cacher, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()

	return &Cacher{client}, err
}

func (c *Cacher) Set(key string, value []byte, expiration int) error {
	return c.redis.Set(key, value, time.Duration(expiration)*time.Second).Err()
}

func (c *Cacher) Get(key string) ([]byte, error) {
	return c.redis.Get(key).Bytes()
}
