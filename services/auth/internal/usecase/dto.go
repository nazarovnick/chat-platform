package usecase

import (
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/session"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/token"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	"time"
)

// LoginInput holds the credentials and metadata required for login.
type LoginInput struct {
	Login     user.Login        // User login (e.g., email or username).
	Password  user.Password     // User password.
	IP        session.IPAddress // IP address of the login request.
	UserAgent session.UserAgent // User agent string from the client.
}

// LoginOutput contains tokens and expiration times returned after a successful login.
type LoginOutput struct {
	AccessToken         token.AccessToken  // Access token for authenticated requests.
	AccessTokenExpires  time.Time          // Expiration time of the access token.
	RefreshToken        token.RefreshToken // Refresh token for session continuation.
	RefreshTokenExpires time.Time          // Expiration time of the refresh token.
}

// LogoutInput is used to log out from the current session.
type LogoutInput struct {
	SessionID session.SessionID // ID of the session to be terminated.
	UserID    user.UserID       // ID of the user performing the logout.
}

// LogoutOutput indicates the result of a single-session logout operation.
type LogoutOutput struct {
	Success bool // True if the logout was successful.
}

// LogoutAllInput is used to log out from all devices, including the current one
type LogoutAllInput struct {
	SessionID session.SessionID // ID of the session requesting the logout.
	UserID    user.UserID       // ID of the user who wants to log out everywhere.
}

// LogoutAllOutput tells if logout from all devices was successful.
type LogoutAllOutput struct {
	Success bool // True if all sessions were logged out.
}

// RefreshInput is used to request new tokens using an existing refresh token.
type RefreshInput struct {
	RefreshToken token.RefreshToken // The old refresh token from the client.
	IP           session.IPAddress  // IP address where the request comes from.
	UserAgent    session.UserAgent  // Browser or device info from the client.
}

// RefreshOutput contains new access and refresh tokens with their expiration times.
type RefreshOutput struct {
	AccessToken         token.AccessToken  // Access token for authenticated requests.
	AccessTokenExpires  time.Time          // Expiration time of the access token.
	RefreshToken        token.RefreshToken // Refresh token for session continuation.
	RefreshTokenExpires time.Time          // Expiration time of the refresh token.
}

// RevokeSessionInput is used to log out from a specific session (device),
// typically triggered from a user session management UI.
type RevokeSessionInput struct {
	SessionID session.SessionID // ID of the session to revoke.
	UserID    user.UserID       // ID of the user who owns the session.
}

// RevokeSessionOutput tells if the session was successfully revoked.
type RevokeSessionOutput struct {
	Success bool // True if the session was revoked successfully.
}
