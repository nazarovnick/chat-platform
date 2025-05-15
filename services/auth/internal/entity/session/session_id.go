package session

import (
	"github.com/google/uuid"
	"strings"
)

// SessionID uniquely identifies a session.
type SessionID uuid.UUID

// NewSessionID generates a new SessionID.
func NewSessionID() SessionID {
	return SessionID(uuid.New())
}

// Validate checks if the SessionID is a valid, non-zero UUID.
func (id SessionID) Validate() error {
	_, err := uuid.Parse(id.String())
	if err != nil {
		return ErrInvalidSessionID
	}
	if uuid.UUID(id) == uuid.Nil {
		return ErrEmptySessionID
	}
	return nil
}

// String returns the string representation of the SessionID.
func (id SessionID) String() string {
	return uuid.UUID(id).String()
}

// ParseSessionID parses and validates a session ID from string.
func ParseSessionID(v string) (SessionID, error) {
	v = strings.TrimSpace(v)
	if v == "" {
		return SessionID{}, ErrEmptySessionID
	}

	id, err := uuid.Parse(v)
	if err != nil {
		return SessionID{}, ErrInvalidSessionID
	}

	return SessionID(id), nil
}
