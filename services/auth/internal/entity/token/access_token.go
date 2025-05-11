package token

import (
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	"time"
)

// AccessToken represents a short-lived access token string.
type AccessToken string

// NewAccessToken creates a new AccessToken from a raw string.
func NewAccessToken(value string) AccessToken {
	return AccessToken(value)
}

// String hides the token value from fmt.Print and %v.
func (t AccessToken) String() string {
	return "[HIDDEN]"
}

// GoString hides the token value from %#v.
func (t AccessToken) GoString() string {
	return "AccessToken([HIDDEN])"
}

// AccessTokenClaims holds access token payload data.
type AccessTokenClaims struct {
	userID    user.UserID
	role      user.Role
	expiresAt time.Time
	createdAt time.Time
}

// NewAccessTokenClaims creates and validates a new access token claims object.
func NewAccessTokenClaims(userID user.UserID, role user.Role, ttl time.Duration) (*AccessTokenClaims, error) {
	now := time.Now().UTC()
	c := &AccessTokenClaims{
		userID:    userID,
		role:      role,
		createdAt: now,
		expiresAt: now.Add(ttl),
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}

// Validate checks if the claims fields are valid.
func (c *AccessTokenClaims) Validate() error {
	if err := c.userID.Validate(); err != nil {
		return err
	}
	if err := c.role.Validate(); err != nil {
		return err
	}
	if c.expiresAt.IsZero() || c.createdAt.IsZero() {
		return ErrInvalidTokenTimestamps
	}
	return nil
}

// UserID returns the user ID.
func (c *AccessTokenClaims) UserID() user.UserID {
	return c.userID
}

// Role returns the user role.
func (c *AccessTokenClaims) Role() user.Role {
	return c.role
}

// CreatedAt returns the token creation time.
func (c *AccessTokenClaims) CreatedAt() time.Time {
	return c.createdAt
}

// ExpiresAt returns the token expiration time.
func (c *AccessTokenClaims) ExpiresAt() time.Time {
	return c.expiresAt
}
