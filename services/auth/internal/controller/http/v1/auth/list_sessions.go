package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http/errors"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	"github.com/nazarovnick/chat-platform/services/auth/internal/usecase"
)

// ListSessionsHandler returns a list of active sessions for the given user.
//
// @Summary      List user sessions
// @Description  Retrieves all active sessions associated with the provided user ID.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      ListSessionsRequest   true  "User ID to retrieve sessions for"
// @Success      200      {object}  ListSessionsResponse  "List of active sessions"
// @Failure      400      {object}  fiber.Error           "Bad request (validation error)"
// @Failure      500      {object}  fiber.Error           "Internal server error while listing sessions"
// @Router       /api/v1/auth/sessions [post]
func ListSessionsHandler(uc usecase.ListSessionsUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req ListSessionsRequest
		if err := c.BodyParser(&req); err != nil {
			return errors.ErrBadRequestBody
		}

		userID, err := user.ParseUserID(req.UserID)
		if err != nil {
			return err
		}

		input := &usecase.ListSessionsInput{UserID: userID}

		output, err := uc.Execute(c.Context(), input)
		if err != nil {
			return err
		}

		resp := &ListSessionsResponse{Sessions: MapSessionList(output.Sessions)}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}
