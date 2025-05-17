package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http/errors"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/session"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	"github.com/nazarovnick/chat-platform/services/auth/internal/usecase"
)

// RevokeSessionHandler handles revocation of a user session by ID.
//
// @Summary      Revoke user session
// @Description  Invalidates a specific session if it belongs to the requesting user.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      RevokeSessionRequest   true  "Session and User ID to revoke"
// @Success      200      {object}  RevokeSessionResponse  "Session revoked successfully"
// @Failure      400      {object}  fiber.Error            "Invalid request data"
// @Failure      403      {object}  fiber.Error            "Access denied to the session"
// @Failure      404      {object}  fiber.Error            "Session not found"
// @Failure      500      {object}  fiber.Error            "Failed to invalidate session"
// @Router       /api/v1/auth/sessions/revoke [post]
func RevokeSessionHandler(uc usecase.RevokeSessionUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req RevokeSessionRequest
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

		input := &usecase.RevokeSessionInput{
			SessionID: sessionID,
			UserID:    userID,
		}

		output, err := uc.Execute(c.Context(), input)
		if err != nil {
			return err
		}
		resp := &RevokeSessionResponse{Success: output.Success}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}
