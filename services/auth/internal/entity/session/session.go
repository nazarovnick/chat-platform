package session

import (
	"errors"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	"time"
)

// Session represents a user session associated with a refresh token.
type Session struct {
	id               SessionID
	userID           user.UserID
	refreshTokenHash RefreshTokenHash
	userAgent        UserAgent
	ip               IPAddress
	createdAt        time.Time
	updatedAt        time.Time
	expiresAt        time.Time
	revoked          bool
}

// NewSession creates and returns a new Session instance.
func NewSession(
	id SessionID,
	userID user.UserID,
	refreshTokenHash RefreshTokenHash,
	userAgent UserAgent,
	ip IPAddress,
	createdAt, updatedAt, expiresAt time.Time,
) (*Session, error) {
	s := &Session{
		id:               id,
		userID:           userID,
		refreshTokenHash: refreshTokenHash,
		userAgent:        userAgent,
		ip:               ip,
		createdAt:        createdAt,
		updatedAt:        updatedAt,
		expiresAt:        expiresAt,
		revoked:          false,
	}
	if err := s.Validate(); err != nil {
		return nil, err
	}

	return s, nil
}

// Validate checks that the session fields are valid.
func (s *Session) Validate() error {
	if s.userID == (user.UserID{}) {
		return errors.New("session: userID is required")
	}
	if err := s.refreshTokenHash.Validate(); err != nil {
		return err
	}
	if err := s.userAgent.Validate(); err != nil {
		return err
	}
	if err := s.ip.Validate(); err != nil {
		return err
	}
	if s.createdAt.IsZero() {
		return ErrEmptyCreatedAt
	}
	if s.updatedAt.IsZero() {
		return ErrEmptyUpdatedAt
	}
	if s.expiresAt.IsZero() {
		return ErrEmptyExpiresAt
	}
	return nil
}

// ID returns the session ID. Getter.
func (s *Session) ID() SessionID {
	return s.id
}

// UserID returns the ID of the user who owns the session. Getter.
func (s *Session) UserID() user.UserID {
	return s.userID
}

// RefreshTokenHash returns the hashed refresh token. Getter.
func (s *Session) RefreshTokenHash() RefreshTokenHash {
	return s.refreshTokenHash
}

// UserAgent returns the user agent string. Getter.
func (s *Session) UserAgent() UserAgent {
	return s.userAgent
}

// IP returns the IP address. Getter.
func (s *Session) IP() IPAddress {
	return s.ip
}

// CreatedAt returns when the session was created. Getter.
func (s *Session) CreatedAt() time.Time {
	return s.createdAt
}

// ExpiresAt returns when the session expires. Getter.
func (s *Session) UpdatedAt() time.Time {
	return s.updatedAt
}

// UpdatedAt returns when the session was last updated. Getter.
func (s *Session) ExpiresAt() time.Time {
	return s.expiresAt
}

// IsRevoked returns whether the session has been revoked. Getter.
func (s *Session) Revoked() bool {
	return s.revoked
}

// SetRefreshTokenHash updates the hashed refresh token.
func (s *Session) SetRefreshTokenHash(hash RefreshTokenHash) {
	s.refreshTokenHash = hash
	s.updatedAt = time.Now()
}

// SetUserAgent updates the user agent string.
func (s *Session) SetUserAgent(agent UserAgent) {
	s.userAgent = agent
	s.updatedAt = time.Now()
}

// SetIPAddress updates the session's IP address.
func (s *Session) SetIPAddress(ip IPAddress) {
	s.ip = ip
	s.updatedAt = time.Now()
}

// SetExpiresAt updates the session expiration time.
func (s *Session) SetExpiresAt(t time.Time) {
	s.expiresAt = t
	s.updatedAt = time.Now()
}

// Revoke marks the session as revoked and updates the timestamp.
func (s *Session) Revoke() {
	s.revoked = true
	s.updatedAt = time.Now()
}

// Unrevoke clears the revoked flag
func (s *Session) Unrevoke() {
	s.revoked = false
	s.updatedAt = time.Now()
}
