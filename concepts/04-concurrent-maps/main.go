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
Expected:
safe execution
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
Expected:
safe + faster reads
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
Q5: Lock Missing (Subtle Race)
========================================
Expected:
race condition
*/
func q5() {
	m := make(map[int]int)
	var mu sync.Mutex

	go func() {
		mu.Lock()
		m[1] = 1
		mu.Unlock()
	}()

	// ❌ read without lock
	fmt.Println(m[1])
}

/*
========================================
Q6: sync.Map Basic Usage
========================================
Expected:
1 true
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
Q8: sync.Map vs map Trap
========================================
Expected:
discussion
*/
func q8() {
	// When to use sync.Map?
	// high read-heavy workloads
}

/*
========================================
Q9: Double Check Locking Pattern
========================================
Expected:
safe lazy init
*/
func q9() {
	m := make(map[string]int)
	var mu sync.Mutex

	key := "a"

	mu.Lock()
	if _, ok := m[key]; !ok {
		m[key] = 1
	}
	mu.Unlock()

	fmt.Println(m[key])
}

/*
========================================
Q10: Goroutine + Map + WaitGroup (REAL)
========================================
Expected:
100 elements
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
