package concurrencypatterns

import (
	"context"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	got := WorkerPool(context.Background(), 2, []int{1, 2, 3}, func(_ context.Context, value int) int {
		return value * value
	})
	for index, want := range []int{1, 4, 9} {
		if got[index] != want {
			t.Fatalf("results[%d] = %d, want %d", index, got[index], want)
		}
	}
}

func TestCachesAndRateLimiter(t *testing.T) {
	ttl := NewTTLCache[string, int]()
	ttl.Set("answer", 42, time.Millisecond)
	if got, ok := ttl.Get("answer"); !ok || got != 42 {
		t.Fatalf("TTL cache Get() = %d, %v", got, ok)
	}
	time.Sleep(2 * time.Millisecond)
	if _, ok := ttl.Get("answer"); ok {
		t.Fatal("expired cache value was returned")
	}

	lru := NewLRUCache[string, int](2)
	lru.Set("a", 1)
	lru.Set("b", 2)
	_, _ = lru.Get("a") // a is now most recently used
	lru.Set("c", 3)
	if _, ok := lru.Get("b"); ok {
		t.Fatal("least recently used entry was not evicted")
	}

	limiter := NewRateLimiter(1, time.Hour)
	defer limiter.Stop()
	if !limiter.Allow() || limiter.Allow() {
		t.Fatal("rate limiter did not enforce burst size")
	}
}

func TestSemaphoreRespectsCancellation(t *testing.T) {
	semaphore := NewSemaphore(1)
	if err := semaphore.Acquire(context.Background()); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := semaphore.Acquire(ctx); err == nil {
		t.Fatal("Acquire() succeeded after context cancellation")
	}
	semaphore.Release()
}
