package errors

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrBadRequestBody            = fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	ErrEmptySessionID            = fiber.NewError(fiber.StatusBadRequest, "Session ID cannot be empty")
	ErrInvalidSessionIDFormat    = fiber.NewError(fiber.StatusBadRequest, "Invalid session ID format")
	ErrEmptyUserID               = fiber.NewError(fiber.StatusBadRequest, "User ID cannot be empty")
	ErrInvalidUserID             = fiber.NewError(fiber.StatusBadRequest, "Invalid user ID format")
	ErrSessionNotFound           = fiber.NewError(fiber.StatusNotFound, "Session not found")
	ErrAccessDenied              = fiber.NewError(fiber.StatusForbidden, "Access denied")
	ErrFailedToInvalidateSession = fiber.NewError(fiber.StatusInternalServerError, "Failed to invalidate session")
	ErrInvalidIPAddress          = fiber.NewError(fiber.StatusBadRequest, "Invalid IP address")
	ErrInvalidUserAgent          = fiber.NewError(fiber.StatusBadRequest, "Invalid User-Agent")
	ErrLoginAlreadyExists        = fiber.NewError(fiber.StatusConflict, "Login already exists")
	ErrInvalidRegistrationData   = fiber.NewError(fiber.StatusBadRequest, "Invalid registration data")
	ErrRegisterUserFailed        = fiber.NewError(fiber.StatusInternalServerError, "Failed to register user")
	ErrInvalidLogin              = fiber.NewError(fiber.StatusBadRequest, "Invalid login")
	ErrInvalidPassword           = fiber.NewError(fiber.StatusBadRequest, "Invalid password format. Password must be at least 8 characters long and include a letter, an uppercase letter, a digit, and a special character (e.g. @,!,#, etc.).")
	ErrInvalidRole               = fiber.NewError(fiber.StatusBadRequest, "Invalid role")
	ErrInvalidCredentials        = fiber.NewError(fiber.StatusUnauthorized, "Invalid login or password")
	ErrLogInFailed               = fiber.NewError(fiber.StatusInternalServerError, "Failed to login")
	ErrFailedToListSessions      = fiber.NewError(fiber.StatusInternalServerError, "Failed to list sessions")
)
