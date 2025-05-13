package user

// PasswordHasher defines the contract for hashing a user's password.
type PasswordHasher interface {
	// Hash generates a secure hash from the given plain-text password.
	Hash(Password) (PasswordHash, error)
}

// PasswordVerifier defines an interface for verifying a plain-text password
// against a stored hash.
type PasswordVerifier interface {
	// Verify checks whether the inputPassword matches the stored hash.
	Verify(storedHash string, inputPassword string) error
}

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
