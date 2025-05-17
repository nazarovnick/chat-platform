package auth

import (
	"github.com/nazarovnick/chat-platform/services/auth/internal/usecase"
)

// MapSessionInfo converts a use case layer SessionInfo object
// to a controller layer SessionInfoDTO for API response formatting.
func MapSessionInfo(s *usecase.SessionInfo) *SessionInfoDTO {
	return &SessionInfoDTO{
		SessionID: s.SessionID.String(),
		UserAgent: s.UserAgent.String(),
		IP:        s.IP.String(),
		CreatedAt: s.CreatedAt,
		ExpiresAt: s.ExpiresAt,
	}
}

// MapSessionList maps a slice of SessionInfo objects from the use case layer
// to a slice of SessionInfoDTOs for API response serialization.
func MapSessionList(list []*usecase.SessionInfo) []*SessionInfoDTO {
	out := make([]*SessionInfoDTO, 0, len(list))
	for _, s := range list {
		out = append(out, MapSessionInfo(s))
	}
	return out
}
