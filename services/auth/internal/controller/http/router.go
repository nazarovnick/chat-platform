package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http/errors"
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http/middleware"
	v1 "github.com/nazarovnick/chat-platform/services/auth/internal/controller/http/v1"
)

// NewRouter creates a new Fiber application, groups API routes under specified prefix (e.g. "/api"),
// and registers each version’s routes (e.g., v1, v2, etc.).
func NewRouter() *fiber.App {
	config := FiberConfig()
	errMapper := errors.NewErrorMapper()

	app := fiber.New(config)
	app.Use(cors.New())
	app.Use(middleware.ErrorHandle(errMapper))

	api := app.Group("/api")
	v1.RegisterRoutes(api.Group("/v1"))

	return app
}
