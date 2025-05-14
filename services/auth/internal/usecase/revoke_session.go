package usecase

import (
	"context"
	"github.com/nazarovnick/chat-platform/services/auth/pkg/errors"
)

// revokeSessionUseCase handles the revocation of a specific user session.
type revokeSessionUseCase struct {
	sessions SessionRepo
}

// NewRevokeSessionUseCase creates a new instance of revokeSessionUseCase.
func NewRevokeSessionUseCase(sessions SessionRepo) RevokeSessionUseCase {
	return &revokeSessionUseCase{sessions: sessions}
}

// Execute revokes a specific session if it belongs to the requesting user.
//
// This use case is intended for user-initiated revocation of another session,
// for example, logging out from a different device.
//
// Steps:
//
//  1. Retrieves the session by its ID.
//
//  2. Verifies that the session belongs to the requesting user.
//
//  3. Invalidates the session so its refresh token can no longer be used.
//
//     Returns success confirmation if the session was revoked.
//
// Or returns an error if the session is not found, doesn't belong to the user,
// or if the invalidation process fails.
func (uc *revokeSessionUseCase) Execute(ctx context.Context, in *RevokeSessionInput) (_ *RevokeSessionOutput, err error) {
	const op = "usecase.revokeSessionUseCase.Execute"
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

	// Step 3: Invalidate session
	if err := uc.sessions.Invalidate(ctx, sess.ID()); err != nil {
		return nil, ErrInvalidatingSession
	}

	return &RevokeSessionOutput{Success: true}, nil
}
