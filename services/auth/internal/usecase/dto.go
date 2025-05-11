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
