package user

import "unicode"

// Password represents a validated plain-text password.
type Password string

// NewPassword creates a new Password and validates its constraints.
func NewPassword(raw string) (Password, error) {
	p := Password(raw)
	if err := p.Validate(); err != nil {
		return "", err
	}
	return p, nil
}

// Validate checks whether the password meets security requirements.
func (p Password) Validate() error {
	pass := string(p)

	if len(pass) < 8 {
		return ErrPasswordTooShort
	}
	if !hasLetter(pass) {
		return ErrPasswordNoLetter
	}
	if !hasDigit(pass) {
		return ErrPasswordNoDigit
	}
	if !hasUpper(pass) {
		return ErrPasswordNoUppercase
	}
	if !hasSpecial(pass) {
		return ErrPasswordNoSpecialChar
	}
	return nil
}

// String returns a redacted string for fmt.Print-style logging.
func (p Password) String() string {
	return "[HIDDEN]"
}

// GoString returns a redacted Go-syntax representation.
func (p Password) GoString() string {
	return "Password([HIDDEN])"
}

// Reveal returns the raw password value.
func (p Password) Reveal() string {
	return string(p)
}

// Equal compares two Password values for equality.
func (p Password) Equal(other Password) bool {
	return string(p) == string(other)
}

// hasLetter returns true if the string contains at least one letter.
func hasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// hasDigit returns true if the string contains at least one digit.
func hasDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

// hasUpper returns true if the string contains at least one uppercase letter.
func hasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

// hasSpecial returns true if the string contains at least one special character.
func hasSpecial(s string) bool {
	special := "!@#$%^&*()-_=+[]{}:;\"'<>,.?/"
	for _, r := range s {
		for _, sym := range special {
			if r == sym {
				return true
			}
		}
	}
	return false
}
