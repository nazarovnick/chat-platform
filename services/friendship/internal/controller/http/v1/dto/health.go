package dto

// HealthResponse represents the response structure for the health check endpoint.
// It contains the status of the service (e.g., "ok" if the service is healthy).
type HealthResponse struct {
	Status string `json:"status"`
}
