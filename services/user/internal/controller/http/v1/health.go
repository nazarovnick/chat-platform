package v1

import "github.com/gofiber/fiber/v2"

// RegisterHealthRoutes registers HTTP endpoints for Kubernetes liveness and readiness probes
// under the "/healthz" path on the provided router.
func RegisterHealthRoutes(router fiber.Router) {
	health := router.Group("/healthz")
	health.Get("/liveness", livenessHandler)
	health.Get("/readiness", readinessHandler)
}

type healthResponse struct {
	Status string `json:"status"`
}

func livenessHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(healthResponse{Status: "alive"})
}

func readinessHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(healthResponse{Status: "ready"})
}
