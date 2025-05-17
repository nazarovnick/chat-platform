package user

// Role defines the user's role in the system.
type Role string

const (
	RoleUser      Role = "user"
	RoleAdmin     Role = "admin"
	RoleModerator Role = "moderator"
)

// NewRole creates and validates a new Role.
func NewRole(value string) (Role, error) {
	role := Role(value)
	if err := role.Validate(); err != nil {
		return "", err
	}
	return role, nil
}

// Validate checks if the role is one of the allowed values.
func (r Role) Validate() error {
	switch r {
	case RoleUser, RoleAdmin, RoleModerator:
		return nil
	default:
		return ErrInvalidRole
	}
}

// String returns the string representation of the role.
func (r Role) String() string {
	return string(r)
}
