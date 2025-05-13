package usecase

import "context"

// RevokeSessionUseCase handles the revocation of a specific user session.
type RevokeSessionUseCase struct {
	sessions SessionRepo
}

// NewRevokeSessionUseCase creates a new instance of RevokeSessionUseCase.
func NewRevokeSessionUseCase(sessions SessionRepo) *RevokeSessionUseCase {
	return &RevokeSessionUseCase{sessions: sessions}
}

// Execute revokes a specific session if it belongs to the requesting user.
// It returns an error if the session is not found or does not belong to the user.
func (uc *RevokeSessionUseCase) Execute(ctx context.Context, in *RevokeSessionInput) (*RevokeSessionOutput, error) {

	sess, err := uc.sessions.GetBySessionID(ctx, in.SessionID)
	if err != nil {
		return nil, err
	}

	if sess.UserID() != in.UserID {
		return nil, ErrAccessDenied
	}

	if err := uc.sessions.Invalidate(ctx, sess.ID()); err != nil {
		return nil, err
	}

	return &RevokeSessionOutput{Success: true}, nil
}
