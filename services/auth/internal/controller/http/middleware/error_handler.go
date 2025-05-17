package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResolver interface {
	Resolve(err error) *fiber.Error
}

func ErrorHandle(resolver ErrorResolver) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			return nil
		}
		if fe, ok := err.(*fiber.Error); ok {
			return fe
		}
		return resolver.Resolve(err)
	}
}
