package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/katiamach/cache-service/internal/cacher/memcached"
	"github.com/katiamach/cache-service/internal/service"
)

func main() {
	// cacher, err := redis.New()
	// if err != nil {
	// 	log.Fatal("failed to connect to redis:", err)
	// }

	cacher := memcached.New()

	service := service.New(cacher)

	app := fiber.New()

	app.Get("/:id", service.VerifyCache, service.GetUser)

	app.Listen(":3000")
}
