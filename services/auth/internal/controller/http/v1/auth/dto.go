package auth

// RegisterRequest contains the necessary fields for user registration.
type RegisterRequest struct {
	Name     string `json:"name" example:"Ivan Petrov"`
	Email    string `json:"email" example:"ivan.petrov@example.com"`
	Password string `json:"password" example:"superpassword123"`
}

// RegisterResponse represents the response after a successful user registration.
type RegisterResponse struct {
	ID    string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Email string `json:"email" example:"ivan.petrov@example.com"`
}

// LoginRequest contains the user's credentials for authentication.
type LoginRequest struct {
	Email    string `json:"email" example:"ivan.petrov@example.com"`
	Password string `json:"password" example:"superpassword123"`
}

// LoginResponse contains the access token and expiration time for a logged-in user.
type LoginResponse struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	ExpiresIn   int64  `json:"expires_in" example:"3600"`
}

// LogoutResponse represents the message returned after a user logs out.
type LogoutResponse struct {
	Message string `json:"message" example:"Logout successfully"`
}
