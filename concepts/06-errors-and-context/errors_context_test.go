package errorscontext

import (
	"context"
	"errors"
	"testing"
	"time"
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

func TestStreamStopsOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	stream := Stream(ctx, []int{1, 2, 3})
	if value := <-stream; value != 1 {
		t.Fatalf("first value = %d, want 1", value)
	}
	cancel()

	select {
	case _, open := <-stream:
		if open {
			t.Fatal("stream remained open after cancellation")
		}
	case <-time.After(time.Second):
		t.Fatal("stream did not stop after cancellation")
	}
}
