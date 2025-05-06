package user

import "regexp"

// loginRegexp defines allowed characters and length for a valid login.
var loginRegexp = regexp.MustCompile(`^[a-zA-Z0-9_.]{3,64}$`)

// Login represents a user's login identifier.
type Login string

// String returns the string representation of the login.
func (l Login) String() string {
	return string(l)
}

// Validate checks if the login meets format and length requirements.
func (l Login) Validate() error {
	if l == "" {
		return ErrEmptyLogin
	}
	if len(l) < 3 || len(l) > 64 {
		return ErrInvalidLoginLength
	}
	if !loginRegexp.MatchString(l.String()) {
		return ErrInvalidLoginFormat
	}
	return nil
}

// NewLogin creates a new Login and validates it.
func NewLogin(value string) (Login, error) {
	l := Login(value)
	if err := l.Validate(); err != nil {
		return "", err
	}
	return l, nil
}
