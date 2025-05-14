package usecase

import (
	"context"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/session"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/token"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	"github.com/nazarovnick/chat-platform/services/auth/pkg/errors"
	"time"
)

// refreshUseCase handles logic for refreshing user tokens.
type refreshUseCase struct {
	users    UserRepo
	sessions SessionRepo
	tokens   TokenService
}

// NewRefreshUseCase creates a new instance of refreshUseCase.
func NewRefreshUseCase(sessions SessionRepo, tokens TokenService) RefreshUseCase {
	return &refreshUseCase{sessions: sessions, tokens: tokens}
}

// generateRefreshTokenWithClaims creates a refresh token and its associated claims for the given user ID.
func (uc *refreshUseCase) generateRefreshTokenWithClaims(userID user.UserID) (token.RefreshToken, *token.RefreshTokenClaims, error) {
	rtClaims, err := uc.tokens.GenerateRefreshTokenClaims(userID)
	if err != nil {
		return "", nil, err
	}
	rt, err := uc.tokens.GenerateRefreshToken(rtClaims)
	if err != nil {
		return "", nil, err
	}

	return rt, rtClaims, nil
}

// generateAccessTokenWithClaims creates an access token and its claims for the specified user entity.
func (uc *refreshUseCase) generateAccessTokenWithClaims(u *user.User) (token.AccessToken, *token.AccessTokenClaims, error) {
	atClaims, err := uc.tokens.GenerateAccessTokenClaims(u)
	if err != nil {
		return "", nil, err
	}
	at, err := uc.tokens.GenerateAccessToken(atClaims)
	if err != nil {
		return "", nil, err
	}

	return at, atClaims, nil
}

// createSession creates and stores a new session using the provided input data
// and hashed refresh token. Returns an error if validation or persistence fails.
func (uc *refreshUseCase) createSession(ctx context.Context, userID user.UserID, rt token.RefreshToken, in *RefreshInput) error {
	rtHash, err := uc.tokens.HashRefreshToken(rt)
	if err != nil {
		return err
	}

	now := time.Now()
	sid := session.NewSessionID()
	sess, err := session.NewSession(sid, userID, rtHash, in.UserAgent, in.IP, now, now, now.Add(30*24*time.Hour)) // TODO: Put TTL to config
	if err != nil {
		return err
	}
	if err := sess.Validate(); err != nil {
		return err
	}
	if err := uc.sessions.Create(ctx, sess); err != nil {
		return err
	}
	return nil
}

// Execute handles the token refresh process by validating the refresh token,
// invalidating the old session, creating a new one, and issuing new tokens.
//
// Steps:
//
//  1. Hash the provided refresh token.
//
//  2. Find the session by the hashed token.
//
//  3. Invalidate the old session to prevent reuse.
//
//  4. Generate a new refresh token and its claims.
//
//  5. Create and store a new session using the new refresh token.
//
//  6. Retrieve the user by ID from the session.
//
//  7. Generate a new access token and its claims.
//
//     Returns new access and refresh tokens with their expiration times, or an error if any step fails.
func (uc *refreshUseCase) Execute(ctx context.Context, in *RefreshInput) (_ *RefreshOutput, err error) {
	const op = "usecase.refreshUseCase.Execute"
	defer func() {
		if err != nil {
			err = errors.Wrap(op, err)
		}
	}()

	// Step 1: Hash the provided refresh token.
	hash, err := uc.tokens.HashRefreshToken(in.RefreshToken)
	if err != nil {
		return nil, ErrHashingRefreshToken
	}

	// Step 2: Find the session by the hashed token.
	sess, err := uc.sessions.GetByRefreshToken(ctx, hash)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	// Step 3: Invalidate the old session to prevent reuse.
	if err := uc.sessions.Invalidate(ctx, sess.ID()); err != nil {
		return nil, ErrInvalidatingSession
	}

	// Step 4: Generate a new refresh token and its claims.
	rt, rtClaims, err := uc.generateRefreshTokenWithClaims(sess.UserID())
	if err != nil {
		return nil, ErrGeneratingRefreshToken
	}

	// Step 5: Create and store a new session using the new refresh token.
	if err := uc.createSession(ctx, sess.UserID(), rt, in); err != nil {
		return nil, ErrCreatingSession
	}

	// Step 6: Retrieve the user by ID from the session.
	u, err := uc.users.GetByID(ctx, sess.UserID())
	if err != nil {
		return nil, ErrUserNotFound
	}

	// Step 7: Generate a new access token and its claims.
	at, atClaims, err := uc.generateAccessTokenWithClaims(u)
	if err != nil {
		return nil, ErrGeneratingAccessToken
	}

	return &RefreshOutput{
		AccessToken:         at,
		AccessTokenExpires:  atClaims.ExpiresAt(),
		RefreshToken:        rt,
		RefreshTokenExpires: rtClaims.ExpiresAt(),
	}, nil

}
