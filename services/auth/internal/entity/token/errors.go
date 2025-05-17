package token

import "errors"

var (
	ErrInvalidTokenTimestamps = errors.New("token: timestamps must be set")
)
