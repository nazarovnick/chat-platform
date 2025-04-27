package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http/v1/dto"
)

// RegisterHealthRoutes registers HTTP endpoints for Kubernetes liveness and readiness probes
// under the "/healthz" path on the provided router.
func RegisterHealthRoutes(router fiber.Router) {
	health := router.Group("/healthz")
	health.Get("/liveness", LivenessHandler)
	health.Get("/readiness", ReadinessHandler)
}

// LivenessHandler handles the liveness probe request.
// It returns a 200 OK response with a status indicating the service is alive.
func LivenessHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(dto.HealthResponse{
		Status: "alive"},
	)
}

// ReadinessHandler handles the readiness probe request.
// It returns a 200 OK response with a status indicating the service is ready.
func ReadinessHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(dto.HealthResponse{
		Status: "ready"},
	)
}
