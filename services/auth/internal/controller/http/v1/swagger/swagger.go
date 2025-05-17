package swagger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/nazarovnick/chat-platform/services/auth/docs/swagger"
)

// RegisterSwaggerRoutes registers the Swagger UI handler on the provided router.
func RegisterSwaggerRoutes(router fiber.Router) {
	router.Get("/swagger/*", swagger.HandlerDefault)
}
