package memcached

import (
	"fmt"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
)

// Cacher wraps redis client.
type Cacher struct {
	mc *memcache.Client
}

// New creates new cacher for the service.
func New() *Cacher {
	client := memcache.New(fmt.Sprintf("%s:%s", os.Getenv("MEMCACHED_HOST"), os.Getenv("MEMCACHED_PORT")))

	return &Cacher{client}
}

func (c *Cacher) Set(key string, value []byte, expiration int) error {
	return c.mc.Set(&memcache.Item{Key: key, Value: value, Expiration: 10})
}

func (c *Cacher) Get(key string) ([]byte, error) {
	value, err := c.mc.Get(key)
	if err != nil {
		return nil, err
	}

	return value.Value, nil
}
