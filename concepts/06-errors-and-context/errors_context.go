// Package errorscontext demonstrates error inspection and cancellation-aware work.
package errorscontext

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidEmail = errors.New("invalid email")

type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Reason)
}

func Register(email string) error {
	if !strings.Contains(email, "@") {
		return fmt.Errorf("register user: %w", ErrInvalidEmail)
	}
	return nil
}

func ValidateAge(age int) error {
	if age < 18 {
		return &ValidationError{Field: "age", Reason: "must be at least 18"}
	}
	return nil
}

// Stream stops promptly when ctx is cancelled and always closes its output.
// That contract lets callers stop consuming without leaving the goroutine blocked.
func Stream(ctx context.Context, values []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, value := range values {
			select {
			case out <- value:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}
