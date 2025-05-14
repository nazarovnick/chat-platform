package usecase

import (
	"context"
	"github.com/nazarovnick/chat-platform/services/auth/pkg/errors"
)

// logoutUseCase handles logging out from a single session (device).
type logoutUseCase struct {
	sessions SessionRepo
}

// NewLogoutUseCase creates a new instance of logoutUseCase.
func NewLogoutUseCase(sessions SessionRepo) LogoutUseCase {
	return &logoutUseCase{sessions: sessions}
}

// Execute logs the user out from the current session.
// It invalidates the session so that the refresh token can no longer be used.
//
// This use case is intended for self-initiated logout from the current device.
//
// Steps:
//
//  1. Retrieves the session by ID from the repository.
//
//  2. Verifies that the session belongs to the requesting user.
//
//  3. Invalidates the session to prevent further use of its refresh token.
//
//     Returns success confirmation if the session was successfully invalidated.
//
// Or returns an error if the session is not found, doesn't belong to the user,
// or cannot be invalidated.
func (uc *logoutUseCase) Execute(ctx context.Context, in *LogoutInput) (_ *LogoutOutput, err error) {
	const op = "usecase.logoutUseCase.Execute"
	defer func() {
		if err != nil {
			err = errors.Wrap(op, err)
		}
	}()

	// Step 1: Find the session by ID
	sess, err := uc.sessions.GetBySessionID(ctx, in.SessionID)
	if err != nil {
		return nil, ErrSessionNotFound
	}

	// Step 2: Check that the session belongs to the requesting user
	if sess.UserID() != in.UserID {
		return nil, ErrAccessDenied
	}

	// Step 3: Invalidate the session
	if err := uc.sessions.Invalidate(ctx, in.SessionID); err != nil {
		return nil, ErrInvalidatingSession
	}
	return &LogoutOutput{Success: true}, nil
}
