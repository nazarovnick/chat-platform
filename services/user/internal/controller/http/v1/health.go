package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazarovnick/chat-platform/services/user/internal/controller/http/v1/dto"
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
//
// @Summary      Kubernetes liveness probe
// @Description  Endpoint for checking if the application is live
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object} dto.HealthResponse
// @Failure      500  {object}  dto.HTTPError
// @Router       /healthz/liveness [get]
func LivenessHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(dto.HealthResponse{
		Status: "alive"},
	)
}

// ReadinessHandler handles the readiness probe request.
// It returns a 200 OK response with a status indicating the service is ready.
//
// @Summary      Kubernetes readiness probe
// @Description  Endpoint for checking if the application is ready
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object} dto.HealthResponse
// @Failure      500  {object}  dto.HTTPError
// @Router       /healthz/readiness [get]
func ReadinessHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(dto.HealthResponse{
		Status: "ready"},
	)
}
