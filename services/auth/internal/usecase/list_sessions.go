package usecase

import (
	"context"
	"github.com/nazarovnick/chat-platform/services/auth/pkg/errors"
)

// listSessionsUseCase handles listing all active sessions for a specific user.
type listSessionsUseCase struct {
	sessions SessionLister
}

// NewListSessionUseCase creates a new instance of listSessionsUseCase.
func NewListSessionUseCase(sessions SessionLister) ListSessionsUseCase {
	return &listSessionsUseCase{sessions: sessions}
}

// Execute retrieves all sessions associated with the given user ID.
//
// Steps:
//
//  1. Queries the session repository for all sessions by user ID.
//
//  2. Maps internal session entities to simplified session info objects.
//
//     Returns a list of sessions or an error if the operation fails.
func (uc *listSessionsUseCase) Execute(ctx context.Context, in *ListSessionsInput) (_ *ListSessionsOutput, err error) {
	const op = "usecase.listSessionsUseCase.Execute"
	defer func() {
		if err != nil {
			err = errors.Wrap(op, err)
		}
	}()

	sessList, err := uc.sessions.ListByUser(ctx, in.UserID)
	if err != nil {
		return nil, ErrSessionListingFailed
	}

	out := make([]*SessionInfo, 0, len(sessList))
	for _, s := range sessList {
		out = append(out, &SessionInfo{
			SessionID: s.ID(),
			UserAgent: s.UserAgent(),
			IP:        s.IP(),
			CreatedAt: s.CreatedAt(),
			ExpiresAt: s.ExpiresAt(),
		})
	}
	return &ListSessionsOutput{Sessions: out}, nil
}
