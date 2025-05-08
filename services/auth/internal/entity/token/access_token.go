package token

import (
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	"time"
)

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
	token := &AccessTokenClaims{
		userID:    userID,
		role:      role,
		createdAt: now,
		expiresAt: now.Add(ttl),
	}
	if err := token.Validate(); err != nil {
		return nil, err
	}
	return token, nil
}

// Validate checks if the claims fields are valid.
func (c *AccessTokenClaims) Validate() error {
	if err := c.userID.Validate(); err != nil {
		return err
	}
	if err := c.role.Validate(); err != nil {
		return err
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
