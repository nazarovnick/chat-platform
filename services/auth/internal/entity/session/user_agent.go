package session

const (
	userAgentMaxLength = 1024
)

// UserAgent represents a validated user agent string.
type UserAgent string

// NewUserAgent creates a new UserAgent and validates it.
func NewUserAgent(value string) (UserAgent, error) {
	agent := UserAgent(value)
	if err := agent.Validate(); err != nil {
		return "", err
	}
	return agent, nil
}

// Validate checks that the UserAgent is not empty and not too long.
func (agent UserAgent) Validate() error {
	if agent == "" {
		return ErrEmptyUserAgent
	}
	if len(agent) > userAgentMaxLength {
		return ErrUserAgentIsTooLong
	}
	return nil
}

// String returns the string representation of the UserAgent.
func (agent UserAgent) String() string {
	return string(agent)
}
