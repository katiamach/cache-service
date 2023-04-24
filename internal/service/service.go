package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Cacher interface {
	Set(key string, value []byte, expiration int) error
	Get(key string) ([]byte, error)
}

type Service struct {
	cacher Cacher
}

func New(cacher Cacher) *Service {
	return &Service{cacher: cacher}
}

func (s *Service) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := http.Get("https://jsonplaceholder.typicode.com/users/" + id)
	if err != nil {
		return fmt.Errorf("failed to get user from source: %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	err = s.cacher.Set(id, body, 10)
	if err != nil {
		return fmt.Errorf("failed to set user in cacher: %w", err)
	}

	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return c.JSON(fiber.Map{"data": user})
}
