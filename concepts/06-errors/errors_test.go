package errors

import (
	"errors"
	"testing"
)

func TestErrorInspection(t *testing.T) {
	if err := Register("arpit.example.com"); !errors.Is(err, ErrInvalidEmail) {
		t.Fatalf("Register() error = %v, want wrapped ErrInvalidEmail", err)
	}

	err := ValidateAge(17)
	var validationError *ValidationError
	if !errors.As(err, &validationError) {
		t.Fatalf("ValidateAge() error = %v, want ValidationError", err)
	}
	if validationError.Field != "age" {
		t.Fatalf("field = %q, want age", validationError.Field)
	}
}

func TestValidInput(t *testing.T) {
	if err := Register("arpit@example.com"); err != nil {
		t.Fatalf("Register() error = %v, want nil", err)
	}
	if err := ValidateAge(18); err != nil {
		t.Fatalf("ValidateAge() error = %v, want nil", err)
	}
}
