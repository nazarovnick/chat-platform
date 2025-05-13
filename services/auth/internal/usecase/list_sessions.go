package usecase

import (
	"context"
)

// ListSessionsUseCase handles listing all active sessions for a specific user.
type ListSessionsUseCase struct {
	sessions SessionLister
}

// NewListSessionUseCase creates a new instance of ListSessionsUseCase.
func NewListSessionUseCase(sessions SessionLister) *ListSessionsUseCase {
	return &ListSessionsUseCase{sessions: sessions}
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
func (uc *ListSessionsUseCase) Execute(ctx context.Context, in *ListSessionsInput) (*ListSessionsOutput, error) {
	sessList, err := uc.sessions.ListByUser(ctx, in.UserID)
	if err != nil {
		return nil, err
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
