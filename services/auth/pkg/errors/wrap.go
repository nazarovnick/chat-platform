package errors

import (
	"fmt"
)

// Wrap - is sugar for fmt.Errorf("%s: %w")
func Wrap(msg string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", msg, err)
}

func WrapWith(marker, cause error) error {
	if cause == nil {
		return marker
	}
	return fmt.Errorf("%w: %v", marker, cause)
}
