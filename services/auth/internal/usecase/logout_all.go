package usecase

import (
	"context"
	pkgerrors "github.com/nazarovnick/chat-platform/services/auth/pkg/errors"
)

// logoutAllUseCase handles logging out from all user sessions across all devices.
type logoutAllUseCase struct {
	sessions SessionRepo
}

// NewLogoutAllUseCase creates a new instance of logoutAllUseCase.
func NewLogoutAllUseCase(sessions SessionRepo) LogoutAllUseCase {
	return &logoutAllUseCase{sessions: sessions}
}

// Execute logs the user out from all sessions (all devices), including the current one.
// It invalidates all sessions linked to the given user ID, so no refresh tokens can be reused.
//
// Steps:
//
//  1. Retrieves the current session by its ID.
//
//  2. Verifies that the session belongs to the requesting user.
//
//  3. Invalidates all sessions associated with the user's ID.
//
//     Prevents reuse of any refresh tokens linked to those sessions.
//
// Or returns an error if the session is not found, doesn't belong to the user,
// or if the invalidation process fails.
func (uc *logoutAllUseCase) Execute(ctx context.Context, in *LogoutAllInput) (_ *LogoutAllOutput, err error) {
	const op = "usecase.logoutAllUseCase.Execute"
	defer func() {
		if err != nil {
			err = pkgerrors.Wrap(op, err)
		}
	}()

	// Step 1: Find the session by ID
	sess, err := uc.sessions.GetBySessionID(ctx, in.SessionID)
	if err != nil {
		return nil, pkgerrors.WrapWith(ErrSessionNotFound, err)
	}

	// Step 2: Check that the session belongs to the requesting user
	if sess.UserID() != in.UserID {
		return nil, ErrAccessDenied
	}

	// Step 3: Invalidate all sessions
	if err := uc.sessions.InvalidateAll(ctx, in.UserID); err != nil {
		return nil, pkgerrors.WrapWith(ErrInvalidatingAllSessions, err)
	}
	return &LogoutAllOutput{Success: true}, nil
}
