package usecase

import (
	"context"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/session"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/token"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	pkgerrors "github.com/nazarovnick/chat-platform/services/auth/pkg/errors"
	"time"
)

// loginUseCase handles user authentication logic.
type loginUseCase struct {
	users    UserRepo
	sessions SessionRepo
	tokens   TokenService
	verifier user.PasswordVerifier
}

// NewLoginUseCase creates a new instance of loginUseCase.
func NewLoginUseCase(users UserRepo, sessions SessionRepo, tokens TokenService, verifier user.PasswordVerifier) LoginUseCase {
	return &loginUseCase{users: users, sessions: sessions, tokens: tokens, verifier: verifier}
}

// generateRefreshTokenWithClaims creates a refresh token and its associated claims for the given user ID.
func (uc *loginUseCase) generateRefreshTokenWithClaims(userID user.UserID) (token.RefreshToken, *token.RefreshTokenClaims, error) {
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
func (uc *loginUseCase) generateAccessTokenWithClaims(u *user.User) (token.AccessToken, *token.AccessTokenClaims, error) {
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

// createSession creates and stores a new session using the provided login input data
// and hashed refresh token. Returns an error if validation or persistence fails.
func (uc *loginUseCase) createSession(ctx context.Context, userID user.UserID, rt token.RefreshToken, in *LoginInput) error {
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

// Execute authenticates the user and initializes a new session.
//
// Steps:
//
//  1. Verifies user credentials using the provided login and password.
//
//  2. Creates a refresh token and stores the session.
//
//  3. Generates an access token for authenticated API usage.
//
//     Returns access and refresh tokens along with their expiration times.
//
// Or returns an error if any step of the process fails.
func (uc *loginUseCase) Execute(ctx context.Context, in *LoginInput) (_ *LoginOutput, err error) {
	const op = "usecase.loginUseCase.Execute"
	defer func() {
		if err != nil {
			err = pkgerrors.Wrap(op, err)
		}
	}()

	// Step 1: Retrieve user by login and validate password.
	u, err := uc.users.GetByLogin(ctx, in.Login)
	if err != nil {
		return nil, pkgerrors.WrapWith(ErrInvalidCredentials, err)
	}
	if err := u.CheckPassword(in.Password.Reveal(), uc.verifier); err != nil {
		return nil, pkgerrors.WrapWith(ErrInvalidCredentials, err)
	}

	// Step 2: Generate refresh token, associated claims, create and stores the session.
	rt, rtClaims, err := uc.generateRefreshTokenWithClaims(u.ID())
	if err != nil {
		return nil, pkgerrors.WrapWith(ErrGeneratingRefreshToken, err)
	}
	if err := uc.createSession(ctx, u.ID(), rt, in); err != nil {
		return nil, pkgerrors.WrapWith(ErrCreatingSession, err)
	}

	// Step 3: Generate access token for authenticated API usage and its claims.
	at, atClaims, err := uc.generateAccessTokenWithClaims(u)
	if err != nil {
		return nil, pkgerrors.WrapWith(ErrGeneratingAccessToken, err)
	}

	// Return both tokens and their expiration timestamps.
	return &LoginOutput{
		AccessToken:         at,
		AccessTokenExpires:  atClaims.ExpiresAt(),
		RefreshToken:        rt,
		RefreshTokenExpires: rtClaims.ExpiresAt(),
	}, nil

}
