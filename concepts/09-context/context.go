// Package context contains small, testable examples of cancellation-aware work.
package context

import "context"

// Stream sends values until it has sent all of them or ctx is cancelled.
// It always closes its output, so callers can safely range over the result.
//
// The send is in a select because a caller may stop receiving before all
// values have been delivered. Without the ctx.Done case, that goroutine could
// remain blocked forever.
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

// Wait returns nil when ready is closed, or the context error when cancellation
// or a deadline happens first. It models a context-aware blocking operation.
func Wait(ctx context.Context, ready <-chan struct{}) error {
	select {
	case <-ready:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
