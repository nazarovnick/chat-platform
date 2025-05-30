package v1

import "github.com/gofiber/fiber/v2"

// RegisterRoutes registers all application-specific API routes for version 1
// on the provided Fiber router.
func RegisterRoutes(router fiber.Router) {
	RegisterSwaggerRoutes(router)
	RegisterHealthRoutes(router)
	RegisterAuthRoutes(router)
}
