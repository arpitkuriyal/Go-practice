// Package errors contains small, testable examples of Go error design.
package errors

import (
	"errors"
	"fmt"
	"strings"
)

// ErrInvalidEmail is a sentinel error. Callers can use errors.Is to handle
// this expected condition even after it has been wrapped with more detail.
var ErrInvalidEmail = errors.New("invalid email")

// ValidationError describes which input field was invalid and why. A caller
// that needs these details can retrieve it with errors.As.
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Reason)
}

// Register validates an email and adds the operation that failed to the
// returned error. The %w verb preserves ErrInvalidEmail in the error chain.
func Register(email string) error {
	if !strings.Contains(email, "@") {
		return fmt.Errorf("register user: %w", ErrInvalidEmail)
	}
	return nil
}

// ValidateAge returns a typed error when the age cannot be accepted.
func ValidateAge(age int) error {
	if age < 18 {
		return &ValidationError{Field: "age", Reason: "must be at least 18"}
	}
	return nil
}
