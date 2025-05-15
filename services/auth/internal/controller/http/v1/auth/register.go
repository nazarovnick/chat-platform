package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/nazarovnick/chat-platform/services/auth/internal/usecase"
)

// RegisterHandler handles user registration requests.
//
// @Summary		Register a new user
// @Description	Register a new user by providing login, password and optional role.
// @Tags		auth
// @Accept		json
// @Produce		json
// @Param		request	body	    RegisterRequest	true	"User registration data"
// @Success     201     {object}    RegisterResponse    "User created successfully"
// @Failure     400     {object}    fiber.Error         "Bad request (validation failed)"
// @Failure     409     {object}    fiber.Error         "Conflict (login already exists)"
// @Failure     500     {object}    fiber.Error         "Internal server error"
// @Router      /api/v1/auth/register [post]
func RegisterHandler(uc usecase.RegisterUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req RegisterRequest
		if err := c.BodyParser(&req); err != nil {
			return ErrInvalidRegistrationData
		}

		input := &usecase.RegisterInput{
			Login:    req.Login,
			Password: req.Password,
			Role:     req.Role,
		}
		output, err := uc.Execute(c.Context(), input)
		if err != nil {
			if errors.Is(err, usecase.ErrLoginAlreadyExists) {
				return ErrLoginAlreadyExists
			}
			if errors.Is(err, usecase.ErrInvalidLogin) {
				return ErrInvalidLogin
			}
			if errors.Is(err, usecase.ErrInvalidPassword) {
				return ErrInvalidPassword
			}
			if errors.Is(err, usecase.ErrInvalidRole) {
				return ErrInvalidRole
			}
			return ErrRegisterUserFailed
		}

		resp := RegisterResponse{
			ID: output.UserID.String(),
		}

		return c.Status(fiber.StatusCreated).JSON(resp)
	}
}
