package auth

import (
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/token"
	"time"
)

// RegisterRequest contains the necessary fields for user registration.
type RegisterRequest struct {
	Login    string `json:"login" example:"IvanPetrov"`
	Password string `json:"password" example:"superpassword123"`
	Role     string `json:"role" example:"user"`
}

// RegisterResponse represents the response after a successful user registration.
type RegisterResponse struct {
	// Registered user ID
	ID string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
}

// LoginRequest contains the user's credentials for authentication.
type LoginRequest struct {
	Login    string `login:"email" example:"IvanPetrov"`
	Password string `json:"password" example:"superpassword123"`
}

// LoginResponse contains the access token and expiration time for a logged-in user.
type LoginResponse struct {
	// JWT access token
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`

	// Access token TTL in seconds
	AccessTokenExpiresIn int64 `json:"access_token_expires_in" example:"3600"`

	// JWT refresh token
	RefreshToken string `json:"refresh_token" example:"ZXlKaGJHY2lPaUpJVXpJMU5pSXNJbXRwWkN..."`

	// Refresh token TTL in seconds
	RefreshTokenExpiresIn int64 `json:"refresh_token_expires_in" example:"2592000"`
}

// LogoutRequest specifies which session and user should be logged out.
type LogoutRequest struct {
	// Current session ID
	SessionID string `json:"session_id" example:"1f2c0e30-8a2b-4a3f-b14c-02ac56d3b773"`

	// Authenticated user ID
	UserID string `json:"user_id" example:"84a778c5-1493-4f4b-8057-108ec241cc33"`
}

// LogoutResponse indicates whether logout was successful.
type LogoutResponse struct {
	// True if session was invalidated
	Success bool `json:"success" example:"true"`
}

// LogoutAllRequest contains data to invalidate all sessions for a user.
type LogoutAllRequest struct {
	// Current session ID
	SessionID string `json:"session_id" example:"1f2c0e30-8a2b-4a3f-b14c-02ac56d3b773"`

	// User ID
	UserID string `json:"user_id" example:"84a778c5-1493-4f4b-8057-108ec241cc33"`
}

// LogoutAllResponse indicates the result of logging out from all sessions.
type LogoutAllResponse struct {
	// True if all sessions were invalidated
	Success bool `json:"success" example:"true"`
}

// RefreshRequest contains a refresh token used to generate a new token pair.
type RefreshRequest struct {
	// Refresh token to validate and rotate
	RefreshToken token.RefreshToken
}

// RefreshResponse returns newly issued access and refresh tokens with expirations.
type RefreshResponse struct {
	// New JWT access token
	AccessToken string `json:"access_token" example:"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9..."`

	// Access token TTL in seconds
	AccessTokenExpiresIn int64 `json:"access_token_expires_in" example:"3600"`

	// New JWT refresh token
	RefreshToken string `json:"refresh_token" example:"ZXlKaGJHY2lPaUpJVXpJMU5pSXNJbXRwWkN..."`

	// Refresh token TTL in seconds
	RefreshTokenExpiresIn int64 `json:"refresh_token_expires_in" example:"2592000"`
}

// SessionInfoDTO represents basic information about a user session.
type SessionInfoDTO struct {
	// Session identifier
	SessionID string `json:"session_id" example:"de305d54-75b4-431b-adb2-eb6b9e546014"`

	// Client's user agent
	UserAgent string `json:"user_agent" example:"Mozilla/5.0 (Windows NT 10.0; Win64; x64)"`

	// IP address used during login
	IP string `json:"ip" example:"192.168.1.1"`

	// Session creation timestamp
	CreatedAt time.Time `json:"created_at" example:"2024-01-15T13:45:30.9408589+03:00"`

	// Session expiration timestamp
	ExpiresAt time.Time `json:"expires_at" example:"2024-02-15T13:45:30.9408589+03:00"`
}

// ListSessionsRequest contains the user ID for retrieving active sessions.
type ListSessionsRequest struct {
	// User whose sessions are being queried
	UserID string `json:"user_id" example:"b4a78b51-03d4-456a-8b91-e5db45ff78f1"`
}

// ListSessionsResponse returns a list of active sessions for the user.
type ListSessionsResponse struct {
	// List of active sessions
	Sessions []*SessionInfoDTO `json:"sessions"`
}

// RevokeSessionRequest represents data for revoking a specific user session.
type RevokeSessionRequest struct {
	// Session to revoke
	SessionID string `json:"session_id" example:"1f2c0e30-8a2b-4a3f-b14c-02ac56d3b773"`

	// Owner of the session
	UserID string `json:"user_id" example:"84a778c5-1493-4f4b-8057-108ec241cc33"`
}

// RevokeSessionResponse indicates whether the revocation was successful.
type RevokeSessionResponse struct {
	// True if the session was revoked
	Success bool `json:"success" example:"true"`
}
