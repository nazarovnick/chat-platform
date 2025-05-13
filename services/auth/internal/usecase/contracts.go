package usecase

import (
	"context"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/session"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/token"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
)

// UserRepo defines the interface for user data access.
type UserRepo interface {
	// GetByLogin returns a user by login or an error if not found.
	GetByLogin(context.Context, user.Login) (*user.User, error)

	// GetByID returns a user by UserID or an error if not found.
	GetByID(context.Context, user.UserID) (*user.User, error)

	// Create stores a new user in the storage system.
	Create(context.Context, *user.User) error
}

// SessionRepo defines the interface for session data management.
type SessionRepo interface {
	// Create stores a new session in the repository.
	Create(context.Context, *session.Session) error

	// GetBySessionID returns a session by SessionID.
	GetBySessionID(context.Context, session.SessionID) (*session.Session, error)

	// GetByRefreshToken returns a session by its hashed refresh token.
	GetByRefreshToken(context.Context, session.RefreshTokenHash) (*session.Session, error)

	// Invalidate marks a session as invalid by its ID.
	Invalidate(context.Context, session.SessionID) error

	// InvalidateAll invalidates all sessions for the given user.
	InvalidateAll(context.Context, user.UserID) error
}

// TokenService defines the interface for working with access and refresh tokens.
type TokenService interface {
	// GenerateAccessTokenClaims creates claims for an access token.
	GenerateAccessTokenClaims(*user.User) (*token.AccessTokenClaims, error)

	// GenerateAccessToken creates a signed access token from claims.
	GenerateAccessToken(*token.AccessTokenClaims) (token.AccessToken, error)

	// GenerateRefreshTokenClaims creates claims for a refresh token.
	GenerateRefreshTokenClaims(user.UserID) (*token.RefreshTokenClaims, error)

	// GenerateRefreshToken creates a signed refresh token from claims.
	GenerateRefreshToken(*token.RefreshTokenClaims) (token.RefreshToken, error)

	// HashRefreshToken returns a hashed version of the refresh token.
	HashRefreshToken(token.RefreshToken) (session.RefreshTokenHash, error)
}

// SessionLister defines the interface for retrieving sessions belonging to a specific user.
type SessionLister interface {
	// ListByUser returns all sessions associated with the given user ID.
	ListByUser(context.Context, user.UserID) ([]*session.Session, error)
}
