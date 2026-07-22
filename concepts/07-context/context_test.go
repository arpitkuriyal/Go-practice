package context

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestStreamStopsOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	stream := Stream(ctx, []int{1, 2, 3})
	if value := <-stream; value != 1 {
		t.Fatalf("first value = %d, want 1", value)
	}
	cancel()

	deadline := time.After(time.Second)
	for {
		select {
		case _, open := <-stream:
			if !open {
				return
			}
		case <-deadline:
			t.Fatal("stream did not stop after cancellation")
		}
	}
}

func TestWait(t *testing.T) {
	t.Run("ready first", func(t *testing.T) {
		ready := make(chan struct{})
		close(ready)
		if err := Wait(context.Background(), ready); err != nil {
			t.Fatalf("Wait() error = %v, want nil", err)
		}
	})

	t.Run("context cancelled first", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		err := Wait(ctx, make(chan struct{}))
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("Wait() error = %v, want context.Canceled", err)
		}
	})
}
