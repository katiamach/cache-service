package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/katiamach/cache-service/internal/cacher/redis"
	"github.com/katiamach/cache-service/internal/service"
)

func main() {
	cacher, err := redis.New()
	if err != nil {
		log.Fatal("failed to connect to redis:", err)
	}

	service := service.New(cacher)

	app := fiber.New()

	app.Get("/:id", service.VerifyCache, service.GetUser)

	app.Listen(":3000")
}
