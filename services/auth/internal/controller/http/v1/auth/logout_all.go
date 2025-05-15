package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/session"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	"github.com/nazarovnick/chat-platform/services/auth/internal/usecase"
)

// LogoutAllHandler logs the user out from all devices (sessions).
//
// @Summary		Log out from all sessions
// @Description	Invalidate all active sessions for the user across all devices.
// @Tags		auth
// @Accept		json
// @Produce		json
// @Param		request body     	LogoutAllRequest   true  "Current session ID and user ID"
// @Success		200     {object}    LogoutAllResponse   "All sessions invalidated"
// @Failure		400     {object}    fiber.Error          "Invalid input or request data"
// @Failure		403     {object}    fiber.Error          "Session does not belong to user"
// @Failure		404     {object}    fiber.Error          "Session not found"
// @Failure		500     {object}    fiber.Error          "Internal server error"
// @Router			/api/v1/auth/logout_all [post]
func LogoutAllHandler(uc usecase.LogoutAllUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req LogoutAllRequest
		if err := c.BodyParser(&req); err != nil {
			return ErrBadRequestBody
		}

		sessionID, err := session.ParseSessionID(req.SessionID)
		if err != nil {
			if errors.Is(err, session.ErrEmptySessionID) {
				return ErrEmptySessionID
			}
			return ErrInvalidSessionIDFormat
		}

		userID, err := user.ParseUserID(req.UserID)
		if err != nil {
			if errors.Is(err, user.ErrEmptyUserID) {
				return ErrEmptyUserID
			}
			return ErrInvalidUserID
		}

		input := &usecase.LogoutAllInput{
			SessionID: sessionID,
			UserID:    userID,
		}

		output, err := uc.Execute(c.Context(), input)
		if err != nil {
			if errors.Is(err, usecase.ErrSessionNotFound) {
				return ErrSessionNotFound
			}
			if errors.Is(err, usecase.ErrAccessDenied) {
				return ErrAccessDenied
			}
			if errors.Is(err, usecase.ErrInvalidatingSession) {
				return ErrFailedToInvalidateSession
			}
		}
		resp := &LogoutAllResponse{Success: output.Success}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}
