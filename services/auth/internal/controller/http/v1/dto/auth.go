package dto

// RegisterRequest contains the necessary fields for user registration.
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse represents the response after a successful user registration.
type RegisterResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// LoginRequest contains the user's credentials for authentication.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse contains the access token and expiration time for a logged-in user.
type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// LogoutResponse represents the message returned after a user logs out.
type LogoutResponse struct {
	Message string `json:"message"`
}
