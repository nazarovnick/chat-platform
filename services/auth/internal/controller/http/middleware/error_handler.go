package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorResolver defines an interface for resolving application errors
// to HTTP-specific *fiber.Error values.
type ErrorResolver interface {
	Resolve(err error) *fiber.Error
}

// ErrorHandle returns a Fiber middleware that intercepts handler errors,
// attempts to resolve them using the provided ErrorResolver, and converts
// them into appropriate HTTP responses.
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
