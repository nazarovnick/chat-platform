package user

// PasswordHash represents a hashed user password.
type PasswordHash string

// NewPasswordHash creates a new PasswordHash and validates it.
func NewPasswordHash(value string) (PasswordHash, error) {
	hash := PasswordHash(value)
	if err := hash.Validate(); err != nil {
		return "", err
	}
	return hash, nil
}

// Validate checks if the password hash is not empty.
func (h PasswordHash) Validate() error {
	if h == "" {
		return ErrEmptyPasswordHash
	}
	return nil
}

// String returns the string representation of the password hash.
func (h PasswordHash) String() string {
	return string(h)
}
