package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/session"
	"github.com/nazarovnick/chat-platform/services/auth/internal/usecase"
)

// RefreshHandler handles token refreshing.
//
// @Summary      Refresh access and refresh tokens
// @Description  Validates refresh token and issues new token pair
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body     RefreshRequest   true  "Refresh token"
// @Success      200     {object} RefreshResponse  "New tokens issued"
// @Failure      400     {object} fiber.Error      "Invalid request data"
// @Failure      401     {object} fiber.Error      "Invalid or expired token"
// @Failure      500     {object} fiber.Error      "Internal server error"
// @Router       /api/v1/auth/refresh [post]
func RefreshHandler(uc usecase.RefreshUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req RefreshRequest
		if err := c.BodyParser(&req); err != nil {
			return ErrBadRequestBody
		}

		ip := c.IP()
		ipAddr, err := session.NewIPAddress(ip)
		if err != nil {
			return ErrInvalidIPAddress
		}

		ua := c.Get("User-Agent")
		userAgent, err := session.NewUserAgent(ua)
		if err != nil {
			return ErrInvalidUserAgent
		}

		input := &usecase.RefreshInput{
			RefreshToken: req.RefreshToken,
			IP:           ipAddr,
			UserAgent:    userAgent,
		}

		output, err := uc.Execute(c.Context(), input)
		if err != nil {
			return ErrInvalidRefreshToken
		}

		resp := RefreshResponse{
			AccessToken:           output.AccessToken.String(),
			AccessTokenExpiresIn:  output.AccessTokenExpires.Unix(),
			RefreshToken:          output.RefreshToken.String(),
			RefreshTokenExpiresIn: output.RefreshTokenExpires.Unix(),
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	}
}
