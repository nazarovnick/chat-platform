package usecase

import "errors"

var (
	ErrAccessDenied       = errors.New("usecase: revoke: access denied, forbidden")
	ErrLoginAlreadyExists = errors.New("usecase: register: login already exists")
)
