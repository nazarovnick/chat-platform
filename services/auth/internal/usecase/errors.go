package usecase

import "errors"

var (
	ErrAccessDenied = errors.New("access denied: forbidden")
)
