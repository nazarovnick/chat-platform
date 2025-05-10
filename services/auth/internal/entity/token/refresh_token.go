package token

import (
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/session"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	"time"
)

// RefreshTokenClaims holds the payload data for a refresh token.
type RefreshTokenClaims struct {
	sessionID session.SessionID
	userID    user.UserID
	createdAt time.Time
	expiresAt time.Time
}

// NewRefreshTokenClaims creates a new RefreshTokenClaims with the given session ID, user ID, and TTL.
func NewRefreshTokenClaims(
	sessionID session.SessionID,
	userID user.UserID,
	ttl time.Duration,
) (*RefreshTokenClaims, error) {
	now := time.Now().UTC()
	c := &RefreshTokenClaims{
		sessionID: sessionID,
		userID:    userID,
		createdAt: now,
		expiresAt: now.Add(ttl),
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}

// Validate checks that all required fields in the claims are valid.
func (c *RefreshTokenClaims) Validate() error {
	if err := c.sessionID.Validate(); err != nil {
		return err
	}
	if err := c.userID.Validate(); err != nil {
		return err
	}
	if c.expiresAt.IsZero() || c.createdAt.IsZero() {
		return ErrInvalidTokenTimestamps
	}

	return nil
}

// UserID returns the user ID. Getter.
func (c *RefreshTokenClaims) UserID() user.UserID {
	return c.userID
}

// SessionID returns the session ID. Getter.
func (c *RefreshTokenClaims) SessionID() session.SessionID {
	return c.sessionID
}

// CreatedAt returns the token creation time. Getter.
func (c *RefreshTokenClaims) CreatedAt() time.Time {
	return c.createdAt
}

// ExpiresAt returns the token expiration time. Getter.
func (c *RefreshTokenClaims) ExpiresAt() time.Time {
	return c.expiresAt
}
