package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http/errors"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/session"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	"github.com/nazarovnick/chat-platform/services/auth/internal/usecase"
)

// LogoutHandler handles user logout from a specific session.
//
// @Summary		Log out from a session
// @Description	Invalidates the session so that its refresh token can no longer be used.
// @Tags		auth
// @Accept		json
// @Produce		json
// @Param		request	body	    LogoutRequest	true	"Logout session payload"
// @Success     200     {object}    LogoutResponse       "Successfully logged out"
// @Failure     400     {object}    fiber.Error          "Invalid request data"
// @Failure     403     {object}    fiber.Error          "Access denied to the session"
// @Failure     404     {object}    fiber.Error          "Session not found"
// @Failure     500     {object}    fiber.Error          "Failed to invalidate session"
// @Router      /api/v1/auth/logout [post]
func LogoutHandler(uc usecase.LogoutUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req LogoutRequest
		if err := c.BodyParser(&req); err != nil {
			return errors.ErrBadRequestBody
		}

		sessionID, err := session.ParseSessionID(req.SessionID)
		if err != nil {
			return err
		}

		userID, err := user.ParseUserID(req.UserID)
		if err != nil {
			return err
		}

		input := &usecase.LogoutInput{
			SessionID: sessionID,
			UserID:    userID,
		}

		output, err := uc.Execute(c.Context(), input)
		if err != nil {
			return err
		}
		resp := &LogoutResponse{Success: output.Success}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}
