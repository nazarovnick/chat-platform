package http

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/nazarovnick/chat-platform/services/friendship/internal/controller/http/v1"
	"net/http"
)

// NewRouter creates a new Fiber application, groups API routes under specified prefix (e.g. "/api"),
// and registers each version’s routes (e.g., v1, v2, etc.).
func NewRouter(gwMux http.Handler) *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	v1route := api.Group("/v1")
	v1.RegisterRoutes(v1route, gwMux)

	return app
}
