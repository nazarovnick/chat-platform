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
//
//	@Summary		Register a new user
//	@Description	Register a new user with email and password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.RegisterRequest	true	"User registration data"
//	@Success		201		{object}	dto.RegisterResponse
//	@Failure		400		{object}	dto.HTTPError
//	@Failure		409		{object}	dto.HTTPError
//	@Failure		500		{object}	dto.HTTPError
//	@Router			/auth/register [post]
func RegisterHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(dto.RegisterResponse{
		ID:    "mock ID",
		Email: "mock@email.com",
	})
}

// LoginHandler handles user login requests.
//
//	@Summary		Login a user
//	@Description	Authenticate user and return access token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.LoginRequest	true	"User login credentials"
//	@Success		200		{object}	dto.LoginResponse
//	@Failure		400		{object}	dto.HTTPError
//	@Failure		401		{object}	dto.HTTPError
//	@Failure		500		{object}	dto.HTTPError
//	@Router			/auth/login [post]
func LoginHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(dto.LoginResponse{
		AccessToken: "test_eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		ExpiresIn:   3600,
	})
}

// LogoutHandler handles user logout requests.
//
//	@Summary		Logout a user
//	@Description	Invalidate user session or access token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	dto.LogoutResponse
//	@Failure		401		{object}	dto.HTTPError
//	@Failure		500		{object}	dto.HTTPError
//	@Router			/auth/logout [post]
func LogoutHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(dto.LogoutResponse{
		Message: "Logout successfully",
	})
}
