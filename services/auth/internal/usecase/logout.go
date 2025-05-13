package usecase

import "context"

// LogoutUseCase handles logging out from a single session (device).
type LogoutUseCase struct {
	sessions SessionRepo
}

// NewLogoutUseCase creates a new instance of LogoutUseCase.
func NewLogoutUseCase(sessions SessionRepo) *LogoutUseCase {
	return &LogoutUseCase{sessions: sessions}
}

// Execute logs the user out from the specified session (usually the current device).
// It invalidates the session so that the refresh token can no longer be used.
func (uc *LogoutUseCase) Execute(ctx context.Context, in *LogoutInput) (*LogoutOutput, error) {
	if err := uc.sessions.Invalidate(ctx, in.SessionID); err != nil {
		return nil, err
	}
	return &LogoutOutput{Success: true}, nil
}
