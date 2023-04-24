package main

import (
	"log"

	"github.com/katiamach/cache-service/internal/cacher/redis"
)

func main() {
	_, err := redis.New()
	if err != nil {
		log.Fatal("failed to connect to redis:", err)
	}
}
