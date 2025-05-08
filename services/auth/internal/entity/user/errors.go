package user

import "errors"

var (
	// Login errors
	ErrEmptyLogin         = errors.New("user: login cannot be empty")
	ErrInvalidLoginLength = errors.New("user: login must be 3-64 characters")
	ErrInvalidLoginFormat = errors.New("user: login must contain only a-z, A-Z, 0-9, _")

	// PasswordHash errors
	ErrEmptyPasswordHash = errors.New("user: password hash cannot be empty")
)
