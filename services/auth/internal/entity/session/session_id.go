package session

import "github.com/google/uuid"

// SessionID uniquely identifies a session.
type SessionID uuid.UUID

// NewSessionID generates a new SessionID.
func NewSessionID() SessionID {
	return SessionID(uuid.New())
}

// String returns the string representation of the SessionID.
func (id SessionID) String() string {
	return uuid.UUID(id).String()
}
