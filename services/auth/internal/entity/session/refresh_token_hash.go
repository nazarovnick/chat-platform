package session

// RefreshTokenHash represents a hashed refresh token.
type RefreshTokenHash string

// NewRefreshTokenHash creates a new RefreshTokenHash.
func NewRefreshTokenHash(value string) (RefreshTokenHash, error) {
	hash := RefreshTokenHash(value)
	if err := hash.Validate(); err != nil {
		return "", err
	}
	return hash, nil
}

// Validate checks if the refresh token hash is not empty.
func (h RefreshTokenHash) Validate() error {
	if h == "" {
		return ErrEmptyRefreshTokenHash
	}
	return nil
}

// String returns the string representation of the password hash.
func (h RefreshTokenHash) String() string {
	return string(h)
}
