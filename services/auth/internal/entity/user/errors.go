package user

import "errors"

var (
	// Login errors
	ErrEmptyLogin         = errors.New("user: login cannot be empty")
	ErrInvalidLoginLength = errors.New("user: login must be 3-64 characters")
	ErrInvalidLoginFormat = errors.New("user: login must contain only a-z, A-Z, 0-9, _")

	// Password errors
	ErrPasswordTooShort      = errors.New("user: password must be at least 8 characters long")
	ErrPasswordNoLetter      = errors.New("user: password must contain at least one letter")
	ErrPasswordNoDigit       = errors.New("user: password must contain at least one digit")
	ErrPasswordNoUppercase   = errors.New("user: password must contain at least one uppercase letter")
	ErrPasswordNoSpecialChar = errors.New("user: password must contain at least one special character: !@#$%^&* etc.")

	ErrInvalidCredentials = errors.New("user: invalid credentials")

	// PasswordHash errors
	ErrEmptyPasswordHash = errors.New("user: password hash cannot be empty")

	// Role errors
	ErrInvalidRole = errors.New("invalid role")

	// UserID errors
	ErrEmptyUserID   = errors.New("user: ID cannot be empty")
	ErrInvalidUserID = errors.New("user: ID is invalid")
)
