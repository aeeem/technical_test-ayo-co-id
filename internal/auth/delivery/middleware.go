package auth_http

import (
	"github.com/gofiber/fiber/v2"
)

/*
REQUIRED(Any middleware must have this)

For every middleware we need a config.
In config we also need to define a function which allows us to skip the middleware if return true.
By convention it should be named as "Filter" but any other name will work too.
*/
type Config struct {
	Filter      func(c *fiber.Ctx) bool // Required
	AuthHandler AuthHandler
}

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Filter returns true
		_, err := config.AuthHandler.AuthUsecase.CheckToken(c.Get("Authorization"))
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.Next()
	}
}
