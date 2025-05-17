package usecase

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid login or password")
	ErrLoginAlreadyExists = errors.New("login already exists")
	ErrAccessDenied       = errors.New("access denied, forbidden")
	ErrInvalidLogin       = errors.New("invalid login format")
	ErrInvalidPassword    = errors.New("invalid password format (must be at least 8 characters long and include a letter, an uppercase letter, a digit, and a special character (e.g. @,!,#, etc.))")
	ErrInvalidRole        = errors.New("invalid role")
	ErrHashingPassword    = errors.New("failed to hash password")
	ErrUserCreationFailed = errors.New("failed to create user")

	// Token errors
	ErrHashingRefreshToken    = errors.New("failed to hash refresh token")
	ErrInvalidRefreshToken    = errors.New("invalid or expired refresh token")
	ErrGeneratingRefreshToken = errors.New("failed to generate refresh token")
	ErrGeneratingAccessToken  = errors.New("failed to generate access token")

	// Session errors
	ErrSessionNotFound         = errors.New("session not found for provided refresh token")
	ErrInvalidatingSession     = errors.New("failed to invalidate previous session")
	ErrInvalidatingAllSessions = errors.New("failed to invalidate all sessions")
	ErrCreatingSession         = errors.New("failed to create new session")
	ErrSessionListingFailed    = errors.New("failed to list sessions")
	ErrSessionInvalidateFailed = errors.New("failed to invalidate session")

	// User errors
	ErrUserNotFound = errors.New("user not found by session user ID")
)
