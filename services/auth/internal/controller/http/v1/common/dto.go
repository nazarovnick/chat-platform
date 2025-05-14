package common

// HTTPError represents the standard error response returned by the API.
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message" example:"error description"`
}
