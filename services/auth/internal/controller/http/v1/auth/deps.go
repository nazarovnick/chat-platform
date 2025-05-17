package auth

import "github.com/nazarovnick/chat-platform/services/auth/internal/usecase"

// AuthUseCases aggregates all authentication-related use cases
// that are used by HTTP handlers in the auth module.
type AuthUseCases struct {
	Register      usecase.RegisterUseCase      // Handles user registration
	Login         usecase.LoginUseCase         // Handles user authentication
	Logout        usecase.LogoutUseCase        // Logs out from current session
	LogoutAll     usecase.LogoutAllUseCase     // Logs out from all sessions
	Refresh       usecase.RefreshUseCase       // Refreshes access and refresh tokens
	ListSessions  usecase.ListSessionsUseCase  // Lists all active sessions
	RevokeSession usecase.RevokeSessionUseCase // Revokes a specific session
}

// NewAuthUseCases constructs and returns an AuthUseCases instance
// by injecting each individual use case.
func NewAuthUseCases() *AuthUseCases {

	// TODO: Init UseCases

	return &AuthUseCases{
		//Register:      register,
		//Login:         login,
		//Logout:        logout,
		//LogoutAll:     logoutAll,
		//Refresh:       refresh,
		//ListSessions:  listSessions,
		//RevokeSession: revokeSession,
	}
}
