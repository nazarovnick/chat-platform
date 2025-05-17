package http

import "github.com/gofiber/fiber/v2"

func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:       "auth-service",
		CaseSensitive: true,
		StrictRouting: false,
	}
}
