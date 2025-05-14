package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http/v1/auth"
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http/v1/health"
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http/v1/swagger"
)

// RegisterRoutes registers all application-specific API routes for version 1
// on the provided Fiber router.
func RegisterRoutes(router fiber.Router) {
	swagger.RegisterSwaggerRoutes(router)
	health.RegisterHealthRoutes(router)
	auth.RegisterAuthRoutes(router)
}
