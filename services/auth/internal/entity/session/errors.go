package session

import "errors"

var (
	// RefreshTokenHash errors
	ErrEmptyRefreshTokenHash = errors.New("session: refresh token hash cannot be empty")

	// IPAddress errors
	ErrEmptyIpAddress   = errors.New("session: ip address cannot be empty")
	ErrInvalidIpAddress = errors.New("session: invalid IP address")

	// UserAgent errors
	ErrEmptyUserAgent     = errors.New("session: user agent cannot be empty")
	ErrUserAgentIsTooLong = errors.New("session: user agent is too long")

	// Session errors
	ErrEmptyCreatedAt = errors.New("session: createdAt must not be zero")
	ErrEmptyUpdatedAt = errors.New("session: updatedAt must not be zero")
	ErrEmptyExpiresAt = errors.New("session: expiresAt must not be zero")
)
