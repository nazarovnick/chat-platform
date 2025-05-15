package auth

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterAuthRoutes registers authentication-related routes for user operations.
func RegisterAuthRoutes(router fiber.Router, uc *AuthUseCases) {
	auth := router.Group("/auth")

	// Creates a new user account
	auth.Post("/register", RegisterHandler(uc.Register))

	// Authenticates user and issues token pair
	auth.Post("/login", LoginHandler(uc.Login))

	// Logs out from the current session (device)
	auth.Post("/logout", LogoutHandler(uc.Logout))

	// Logs out from all sessions across all devices
	auth.Post("/logout_all", LogoutAllHandler(uc.LogoutAll))

	// Refreshes access and refresh tokens
	auth.Post("/refresh", RefreshHandler(uc.Refresh))

	// Lists all active user sessions
	auth.Post("/sessions", ListSessionsHandler(uc.ListSessions))

	// Revokes a specific session (e.g., logout from another device)
	auth.Post("/sessions/revoke", RevokeSessionHandler(uc.RevokeSession))
}
