package usecase

import (
	"context"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/session"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/token"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
)

// Interface definitions for dependencies (outbound port)
type (
	// UserRepo defines the interface for user data access.
	UserRepo interface {
		// GetByLogin returns a user by login or an error if not found.
		GetByLogin(context.Context, user.Login) (*user.User, error)

		// GetByID returns a user by UserID or an error if not found.
		GetByID(context.Context, user.UserID) (*user.User, error)

		// Create stores a new user in the storage system.
		Create(context.Context, *user.User) error
	}

	// SessionRepo defines the interface for session data management.
	SessionRepo interface {
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
	TokenService interface {
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
	SessionLister interface {
		// ListByUser returns all sessions associated with the given user ID.
		ListByUser(context.Context, user.UserID) ([]*session.Session, error)
	}
)

// Interface definitions for usecases (inbound port)
type (
	// RegisterUseCase handles user registration logic.
	RegisterUseCase interface {
		Execute(context.Context, *RegisterInput) (*RegisterOutput, error)
	}

	// LoginUseCase authenticates a user and creates a new session.
	LoginUseCase interface {
		Execute(context.Context, *LoginInput) (*LoginOutput, error)
	}

	// LogoutUseCase invalidates a single session by ID.
	LogoutUseCase interface {
		Execute(context.Context, *LogoutInput) (*LogoutOutput, error)
	}

	// LogoutAllUseCase invalidates all sessions for a user.
	LogoutAllUseCase interface {
		Execute(context.Context, *LogoutAllInput) (*LogoutAllOutput, error)
	}

	// RefreshUseCase issues new tokens using a refresh token.
	RefreshUseCase interface {
		Execute(context.Context, *RefreshInput) (*RefreshOutput, error)
	}

	// ListSessionsUseCase returns all active sessions for a user.
	ListSessionsUseCase interface {
		Execute(context.Context, *ListSessionsInput) (*ListSessionsOutput, error)
	}

	// RevokeSessionUseCase invalidates a specific session if owned by the user.
	RevokeSessionUseCase interface {
		Execute(context.Context, *RevokeSessionInput) (*RevokeSessionOutput, error)
	}
)
