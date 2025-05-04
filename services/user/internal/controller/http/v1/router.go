package v1

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// RegisterRoutes registers all application-specific API routes for version 1
// on the provided Fiber router.
func RegisterRoutes(router fiber.Router, gwMux http.Handler) {
	RegisterSwaggerRoutes(router)
	RegisterHealthRoutes(router)
	RegisterGrpcGatewayRoutes(router, gwMux)
}
