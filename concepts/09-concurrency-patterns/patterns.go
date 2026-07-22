// Package concurrencypatterns demonstrates bounded concurrent work and a TTL cache.
package concurrencypatterns

import (
	"container/list"
	"context"
	"sync"
	"time"
)

// WorkerPool applies work to each job with at most workers jobs running at once.
func WorkerPool(ctx context.Context, workers int, jobs []int, work func(context.Context, int) int) []int {
	if workers <= 0 {
		return nil
	}
	results := make([]int, len(jobs))
	jobQueue := make(chan int)
	var wg sync.WaitGroup

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for index := range jobQueue {
				select {
				case <-ctx.Done():
					return
				default:
					results[index] = work(ctx, jobs[index])
				}
			}
		}()
	}

	for index := range jobs {
		select {
		case jobQueue <- index:
		case <-ctx.Done():
			close(jobQueue)
			wg.Wait()
			return results
		}
	}
	close(jobQueue)
	wg.Wait()
	return results
}

type cacheEntry[V any] struct {
	value     V
	expiresAt time.Time
}

// TTLCache is a small, safe cache. It removes expired values lazily on Get.
type TTLCache[K comparable, V any] struct {
	mu    sync.RWMutex
	items map[K]cacheEntry[V]
}

func NewTTLCache[K comparable, V any]() *TTLCache[K, V] {
	return &TTLCache[K, V]{items: make(map[K]cacheEntry[V])}
}

func (c *TTLCache[K, V]) Set(key K, value V, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = cacheEntry[V]{value: value, expiresAt: time.Now().Add(ttl)}
}

func (c *TTLCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.items[key]
	if !ok || time.Now().After(entry.expiresAt) {
		if ok {
			delete(c.items, key)
		}
		var zero V
		return zero, false
	}
	return entry.value, true
}

// Semaphore bounds a resource such as parallel requests or file operations.
type Semaphore struct {
	tokens chan struct{}
}

func NewSemaphore(limit int) *Semaphore {
	return &Semaphore{tokens: make(chan struct{}, limit)}
}

func (s *Semaphore) Acquire(ctx context.Context) error {
	select {
	case s.tokens <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Semaphore) Release() { <-s.tokens }

// RateLimiter allows burst requests and refills one token every interval.
// Stop must be called to release the ticker goroutine.
type RateLimiter struct {
	tokens chan struct{}
	stop   chan struct{}
	once   sync.Once
}

func NewRateLimiter(burst int, interval time.Duration) *RateLimiter {
	limiter := &RateLimiter{tokens: make(chan struct{}, burst), stop: make(chan struct{})}
	for range burst {
		limiter.tokens <- struct{}{}
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case limiter.tokens <- struct{}{}:
				default: // bucket already full
				}
			case <-limiter.stop:
				return
			}
		}
	}()
	return limiter
}

func (r *RateLimiter) Allow() bool {
	select {
	case <-r.tokens:
		return true
	default:
		return false
	}
}

func (r *RateLimiter) Stop() { r.once.Do(func() { close(r.stop) }) }

type lruEntry[K comparable, V any] struct {
	key   K
	value V
}

// LRUCache evicts the least recently used item when capacity is reached.
type LRUCache[K comparable, V any] struct {
	mu       sync.Mutex
	capacity int
	items    map[K]*list.Element
	order    *list.List
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{capacity: capacity, items: make(map[K]*list.Element), order: list.New()}
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	element, ok := c.items[key]
	if !ok {
		var zero V
		return zero, false
	}
	c.order.MoveToFront(element)
	return element.Value.(lruEntry[K, V]).value, true
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if element, ok := c.items[key]; ok {
		element.Value = lruEntry[K, V]{key: key, value: value}
		c.order.MoveToFront(element)
		return
	}
	if c.capacity <= 0 {
		return
	}
	element := c.order.PushFront(lruEntry[K, V]{key: key, value: value})
	c.items[key] = element
	if c.order.Len() <= c.capacity {
		return
	}
	oldest := c.order.Back()
	entry := oldest.Value.(lruEntry[K, V])
	delete(c.items, entry.key)
	c.order.Remove(oldest)
}
