package user

import (
	"fmt"
	"time"
)

// User represents a system user with authentication data.
type User struct {
	id           UserID
	login        Login
	passwordHash PasswordHash
	role         Role
	isBlocked    bool
	createdAt    time.Time
}

// NewUser creates a new User instance with the given parameters.
// If createdAt is zero, the current time will be used.
func NewUser(
	id UserID,
	login Login,
	passwordHash PasswordHash,
	role Role,
	isBlocked bool,
	createdAt time.Time,
) *User {
	if createdAt.IsZero() {
		createdAt = time.Now()
	}
	return &User{
		id:           id,
		login:        login,
		passwordHash: passwordHash,
		role:         role,
		isBlocked:    isBlocked,
		createdAt:    createdAt,
	}
}

// ID returns the user's ID.
func (u *User) ID() UserID { return u.id }

// Login returns the user's login.
func (u *User) Login() Login { return u.login }

// PasswordHash returns the user's hashed password.
func (u *User) PasswordHash() PasswordHash { return u.passwordHash }

// IsBlocked returns whether the user is blocked.
func (u *User) IsBlocked() bool { return u.isBlocked }

// CreatedAt returns the user's creation time.
func (u *User) CreatedAt() time.Time { return u.createdAt }

// SetPasswordHash sets the user's password hash.
func (u *User) SetPasswordHash(hash PasswordHash) { u.passwordHash = hash }

// Block marks the user as blocked.
func (u *User) Block() { u.isBlocked = true }

// Unblock marks the user as not blocked.
func (u *User) Unblock() { u.isBlocked = false }

// SetLogin sets the user's login.
func (u *User) SetLogin(login Login) { u.login = login }

// CheckPassword verifies the input password against the stored hash using the provided verifier.
func (u *User) CheckPassword(inputPassword string, verifier PasswordVerifier) error {
	if err := verifier.Verify(u.passwordHash.String(), inputPassword); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidCredentials, err)
	}
	return nil
}
