package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gofiber/fiber/v2"
)

func (s *Service) VerifyCache(c *fiber.Ctx) error {
	id := c.Params("id")

	value, err := s.cacher.Get(id)

	// if errors.Is(err, redis.Nil) {
	// 	return c.Next()
	if errors.Is(err, memcache.ErrCacheMiss) {
		return c.Next()
	} else if err != nil {
		return err
	}

	data := User{}
	err = json.Unmarshal(value, &data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return c.JSON(fiber.Map{"cached": data})
}
