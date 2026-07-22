package main

import (
	"fmt"
	"sync"
	"time"
)

/*
========================================
Q1: Goroutine Exit Before Execution
========================================
Expected Output:
(maybe nothing)

Reason:
The main goroutine exits immediately.
If the program ends first, other goroutines are
terminated before they get a chance to run.
*/
func q1() {
	go fmt.Println("Hello from goroutine")
}

/*
========================================
Q2: Fix with Sleep (Not Ideal)
========================================
Expected Output:
Hello

Reason:
time.Sleep() gives the goroutine time to run,
but it doesn't guarantee completion.
Useful for demos, not production code.
*/
func q2() {
	go fmt.Println("Hello")
	time.Sleep(time.Millisecond)
}

/*
========================================
Q3: WaitGroup Correct Usage
========================================
Expected Output:
Hello

Reason:
sync.WaitGroup waits until all goroutines call Done().
This is the correct way to wait for goroutines to finish.
*/
func q3() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Hello")
	}()

	wg.Wait()
}

/*
========================================
Q3A: WaitGroup with Multiple Goroutines
========================================
Expected Output:
1000

Reason:
WaitGroup waits until all goroutines finish.
Mutex ensures only one goroutine updates
the shared variable at a time.
*/
func q3A() {
	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		sum int
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock()
			sum++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(sum)
}

/*
========================================
Q4: Loop Variables in Modern Go
========================================
Go 1.22+ creates a fresh loop variable for each iteration declared with :=.

Expected Output:
0 1 2 (order may vary)

Reason:
Each iteration gets its own copy of i.
Every goroutine prints its own value.
*/
func q4() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Millisecond)
}

/*
========================================
Q5: Explicit Value Passing
========================================
Expected Output:
0 1 2 (any order)

Reason:
Passing i as a parameter gives each goroutine
its own copy. This works correctly in every
Go version.
*/
func q5() {
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond)
}

/*
========================================
Q6: Deadlock (No Receiver)
========================================
Expected:
fatal error: deadlock

Reason:
Unbuffered channels require both a sender and
a receiver. Here, the send blocks forever,
causing a deadlock.
*/
func q6() {
	ch := make(chan int)
	ch <- 1
}

/*
========================================
Q7: Fix Deadlock with Goroutine
========================================
Expected Output:
1

Reason:
The sender runs in another goroutine while
the main goroutine receives the value.
Both operations complete successfully.
*/
func q7() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	fmt.Println(<-ch)
}

/*
========================================
Q8: Buffered Channel
========================================
Expected Output:
(no block)

Reason:
A buffered channel can store values without
an immediate receiver. The first send succeeds
because the buffer has capacity 1.
*/
func q8() {
	ch := make(chan int, 1)
	ch <- 1
}

/*
========================================
Q9: Range on Channel (Missing Close)
========================================
Expected:
deadlock

Reason:
range keeps receiving until the channel is closed.
Without close(), it waits forever after all values
are received.
*/
func q9() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		// no close
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

/*
========================================
Q10: Fix Range with Close
========================================
Expected Output:
1
2

Reason:
Closing the channel tells range that no more
values will be sent, allowing the loop to exit.
*/
func q10() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

/*
========================================
Q11: Data Race
========================================
Expected:
race condition

Reason:
Multiple goroutines modify x at the same time.
x++ is not atomic, so the final result is
unpredictable.
*/
func q11() {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			x++
		}()
	}

	time.Sleep(time.Millisecond)
	fmt.Println(x)
}

/*
========================================
Q12: Fix Race with Mutex
========================================
Expected Output:
1000

Reason:
A Mutex allows only one goroutine to update x
at a time. WaitGroup ensures all goroutines
finish before printing.
*/
func q12() {
	x := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			x++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(x)
}

func main() {
	q1()
	q2()
	q3()
	q3A()
	q4()
	q5()
	// q6() // deadlock
	q7()
	q8()
	// q9() // deadlock
	q10()
	q11()
	q12()
}
