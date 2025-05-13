package usecase

import "context"

// LogoutAllUseCase handles logging out from all user sessions across all devices.
type LogoutAllUseCase struct {
	sessions SessionRepo
}

// NewLogoutAllUseCase creates a new instance of LogoutAllUseCase.
func NewLogoutAllUseCase(sessions SessionRepo) *LogoutAllUseCase {
	return &LogoutAllUseCase{sessions: sessions}
}

// Execute logs the user out from all sessions (all devices), including the current one.
// It invalidates all sessions linked to the given user ID, so no refresh tokens can be reused.
func (uc *LogoutAllUseCase) Execute(ctx context.Context, in *LogoutAllInput) (*LogoutAllOutput, error) {
	if err := uc.sessions.InvalidateAll(ctx, in.UserID); err != nil {
		return nil, err
	}
	return &LogoutAllOutput{Success: true}, nil
}
