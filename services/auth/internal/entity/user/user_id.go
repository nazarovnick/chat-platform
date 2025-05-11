package user

import "github.com/google/uuid"

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
