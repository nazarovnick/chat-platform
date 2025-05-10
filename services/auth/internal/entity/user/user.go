package user

import (
	"github.com/google/uuid"
	"time"
)

// UserID represents a unique identifier for a user.
type UserID uuid.UUID

// NewUserID generates a new UserID.
func NewUserID() UserID {
	return UserID(uuid.New())
}

// Validate checks whether the UserID is valid and non-zero.
func (id UserID) Validate() error {
	_, err := uuid.Parse(id.String())
	if err != nil {
		return ErrInvalidUserID
	}
	if uuid.UUID(id) == uuid.Nil {
		return ErrEmptyUserID
	}
	return nil
}

// String returns the string representation of the UserID.
func (id UserID) String() string {
	return uuid.UUID(id).String()
}

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

// ID returns the user's ID. Getter.
func (u *User) ID() UserID { return u.id }

// Login returns the user's login. Getter.
func (u *User) Login() Login { return u.login }

// PasswordHash returns the user's hashed password. Getter.
func (u *User) PasswordHash() PasswordHash { return u.passwordHash }

// IsBlocked returns whether the user is blocked. Getter.
func (u *User) IsBlocked() bool { return u.isBlocked }

// CreatedAt returns the user's creation time. Getter.
func (u *User) CreatedAt() time.Time { return u.createdAt }

// SetPasswordHash sets the user's password hash. Setter.
func (u *User) SerPasswordHash(hash PasswordHash) { u.passwordHash = hash }

// Block marks the user as blocked. Setter.
func (u *User) Block() { u.isBlocked = true }

// Unblock marks the user as not blocked. Setter.
func (u *User) Unblock() { u.isBlocked = false }

// SetLogin sets the user's login. Setter.
func (u *User) SetLogin(login Login) { u.login = login }
