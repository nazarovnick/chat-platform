package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http/v1/dto"
)

// RegisterAuthRoutes registers authentication-related routes for user operations.
func RegisterAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Get("/register", RegisterHandler)
	auth.Get("/login", LoginHandler)
	auth.Get("/logout", LogoutHandler)
}

// RegisterHandler handles user registration requests.
func RegisterHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(dto.RegisterResponse{
		ID:    "mock ID",
		Email: "mock@email.com",
	})
}

// LoginHandler handles user login requests.
func LoginHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(dto.LoginResponse{
		AccessToken: "mocktoken",
		ExpiresIn:   123,
	})
}

// LogoutHandler handles user logout requests.
func LogoutHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(dto.LogoutResponse{
		Message: "Logout successfully",
	})
}
