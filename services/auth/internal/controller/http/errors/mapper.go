package errors

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/session"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	"github.com/nazarovnick/chat-platform/services/auth/internal/usecase"
)

// ErrorMapper maps domain and usecase errors to corresponding *fiber.Error responses.
type ErrorMapper map[error]*fiber.Error

// NewErrorMapper returns a new ErrorMapper with predefined mappings
// from domain and usecase errors to HTTP-specific *fiber.Error values.
func NewErrorMapper() ErrorMapper {
	m := ErrorMapper{
		// Session entity errors mapping
		session.ErrEmptyIpAddress:     ErrInvalidIPAddress,
		session.ErrInvalidIpAddress:   ErrInvalidIPAddress,
		session.ErrEmptyUserAgent:     ErrInvalidUserAgent,
		session.ErrUserAgentIsTooLong: ErrInvalidUserAgent,
		session.ErrEmptySessionID:     ErrEmptySessionID,
		session.ErrInvalidSessionID:   ErrInvalidSessionIDFormat,

		// User entity errors mapping
		user.ErrEmptyUserID:           ErrEmptyUserID,
		user.ErrInvalidUserID:         ErrInvalidUserID,
		user.ErrEmptyLogin:            ErrInvalidLogin,
		user.ErrInvalidLoginLength:    ErrInvalidLogin,
		user.ErrInvalidLoginFormat:    ErrInvalidLogin,
		user.ErrPasswordTooShort:      ErrInvalidPassword,
		user.ErrPasswordNoLetter:      ErrInvalidPassword,
		user.ErrPasswordNoDigit:       ErrInvalidPassword,
		user.ErrPasswordNoUppercase:   ErrInvalidPassword,
		user.ErrPasswordNoSpecialChar: ErrInvalidPassword,

		// Usecases errors mapping
		usecase.ErrSessionListingFailed:   ErrFailedToListSessions,
		usecase.ErrInvalidCredentials:     ErrInvalidCredentials,
		usecase.ErrGeneratingRefreshToken: ErrLogInFailed,
		usecase.ErrCreatingSession:        ErrLogInFailed,
		usecase.ErrGeneratingAccessToken:  ErrLogInFailed,
		usecase.ErrSessionNotFound:        ErrSessionNotFound,
		usecase.ErrAccessDenied:           ErrAccessDenied,
		usecase.ErrInvalidatingSession:    ErrFailedToInvalidateSession,
		usecase.ErrHashingRefreshToken:    ErrFailedToInvalidateSession,
		usecase.ErrInvalidRefreshToken:    ErrFailedToInvalidateSession,
		usecase.ErrUserNotFound:           ErrFailedToInvalidateSession,
		usecase.ErrLoginAlreadyExists:     ErrLoginAlreadyExists,
		usecase.ErrInvalidLogin:           ErrInvalidLogin,
		usecase.ErrInvalidPassword:        ErrInvalidPassword,
		usecase.ErrInvalidRole:            ErrInvalidRole,
		usecase.ErrHashingPassword:        ErrRegisterUserFailed,
		usecase.ErrUserCreationFailed:     ErrRegisterUserFailed,
	}

	return m
}

// Resolve returns the corresponding *fiber.Error for a given error.
// If no match is found, fiber.ErrInternalServerError is returned.
func (m ErrorMapper) Resolve(err error) *fiber.Error {
	for target, mapped := range m {
		if errors.Is(err, target) {
			return mapped
		}
	}
	return fiber.ErrInternalServerError
}
