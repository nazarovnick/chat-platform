package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http/errors"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/session"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	"github.com/nazarovnick/chat-platform/services/auth/internal/usecase"
)

// LoginHandler handles user login requests.
//
// @Summary		Log in
// @Description	Authenticate user with login and password. Returns access and refresh tokens.
// @Tags		auth
// @Accept		json
// @Produce		json
// @Param		request body        LoginRequest   true  "User login credentials"
// @Success		200     {object}    LoginResponse      "Tokens issued successfully"
// @Failure		400     {object}    fiber.Error        "Invalid request data or input format"
// @Failure		401     {object}    fiber.Error        "Invalid login or password"
// @Failure		500     {object}    fiber.Error        "Internal server error"
// @Router		/api/v1/auth/login [post]
func LoginHandler(uc usecase.LoginUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return errors.ErrBadRequestBody
		}

		ip := c.IP()
		ipAddr, err := session.NewIPAddress(ip)
		if err != nil {
			return err
		}

		ua := c.Get("User-Agent")
		userAgent, err := session.NewUserAgent(ua)
		if err != nil {
			return err
		}

		login, err := user.NewLogin(req.Login)
		if err != nil {
			return err
		}

		password, err := user.NewPassword(req.Password)
		if err != nil {
			return err
		}

		input := &usecase.LoginInput{
			Login:     login,
			Password:  password,
			IP:        ipAddr,
			UserAgent: userAgent,
		}

		output, err := uc.Execute(c.Context(), input)
		if err != nil {
			return err
		}

		resp := LoginResponse{
			AccessToken:           output.AccessToken.String(),
			AccessTokenExpiresIn:  output.AccessTokenExpires.Unix(),
			RefreshToken:          output.RefreshToken.String(),
			RefreshTokenExpiresIn: output.RefreshTokenExpires.Unix(),
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	}
}
