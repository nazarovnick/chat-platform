package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/nazarovnick/chat-platform/services/friendship/api/openapi/v1"
)

// RegisterSwaggerRoutes registers the Swagger UI handler on the provided router.
func RegisterSwaggerRoutes(router fiber.Router) {
	router.Get("/swagger/*", swagger.HandlerDefault)
}
