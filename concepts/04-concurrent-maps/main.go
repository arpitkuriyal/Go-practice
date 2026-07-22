package main

import (
	"fmt"
	"sync"
)

/*
========================================
Q1: Concurrent Map Write (CRASH)
========================================
Expected:
fatal error: concurrent map writes

Reason:
A regular map is not safe for concurrent writes.
Multiple goroutines writing at the same time
cause a runtime panic.
*/
func q1() {
	m := make(map[int]int)

	for i := 0; i < 100; i++ {
		go func(i int) {
			m[i] = i
		}(i)
	}
}

/*
========================================
Q2: Concurrent Read + Write (CRASH)
========================================
Expected:
fatal error OR race

Reason:
Reading and writing a regular map concurrently
is unsafe. Protect the map with synchronization
or use sync.Map.
*/
func q2() {
	m := make(map[int]int)

	go func() {
		m[1] = 1
	}()

	fmt.Println(m[1])
}

/*
========================================
Q3: Fix with Mutex
========================================
Expected Output:
100

Reason:
A Mutex allows only one goroutine to access
the map at a time, making concurrent writes safe.
*/
func q3() {
	m := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(len(m))
}

/*
========================================
Q4: Read Lock Optimization (RWMutex)
========================================
Expected Output:
1

Reason:
RWMutex allows multiple readers at the same time,
but only one writer. Use RLock() for reads and
Lock() for writes.
*/
func q4() {
	m := make(map[int]int)
	var mu sync.RWMutex

	mu.Lock()
	m[1] = 1
	mu.Unlock()

	mu.RLock()
	fmt.Println(m[1])
	mu.RUnlock()
}

/*
========================================
Q5: Missing Read Lock
========================================
Expected:
race condition

Reason:
Locking only writes is not enough.
Reads must also be protected when writes
can happen concurrently.
*/
func q5() {
	m := make(map[int]int)
	var mu sync.Mutex

	go func() {
		mu.Lock()
		m[1] = 1
		mu.Unlock()
	}()

	// ❌ Read without lock
	fmt.Println(m[1])
}

/*
========================================
Q6: sync.Map Basic Usage
========================================
Expected Output:
1 true

Reason:
sync.Map is safe for concurrent access.
Store() adds a value and Load() retrieves it.
*/
func q6() {
	var m sync.Map

	m.Store("a", 1)

	val, ok := m.Load("a")
	fmt.Println(val, ok)
}

/*
========================================
Q7: sync.Map Range
========================================
Expected:
prints key-value

Reason:
Range() iterates over every key-value pair
stored inside a sync.Map.
*/
func q7() {
	var m sync.Map

	m.Store("a", 1)
	m.Store("b", 2)

	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}

/*
========================================
Q8: sync.Map vs map
========================================
Expected:
discussion

Reason:
Use map + Mutex for most applications.
Use sync.Map mainly for read-heavy workloads
or independent concurrent keys.
*/
func q8() {
	// When to use sync.Map?
	// High read-heavy workloads.
}

/*
========================================
Q9: sync.Once (Run Only Once)
========================================
Expected Output:
Initializing...
Done

Reason:
sync.Once guarantees a function executes only once,
even if multiple goroutines call it simultaneously.
Useful for lazy initialization.
*/
func q9() {
	var once sync.Once
	var wg sync.WaitGroup

	init := func() {
		fmt.Println("Initializing...")
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			once.Do(init)
		}()
	}

	wg.Wait()
	fmt.Println("Done")
}

/*
========================================
Q10: WaitGroup + Mutex (Real Example)
========================================
Expected Output:
100

Reason:
WaitGroup waits for all goroutines to finish,
while Mutex protects the shared map from
concurrent writes.
*/
func q10() {
	m := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(len(m))
}

func main() {
	// q1() // crash
	// q2() // crash
	q3()
	q4()
	// q5() // race
	q6()
	q7()
	q8()
	q9()
	q10()
}
