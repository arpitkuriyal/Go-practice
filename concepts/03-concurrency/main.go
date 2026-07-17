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
Q4: Loop Variable Capture Trap
========================================
Expected Output:
3 3 3 (order may vary)
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
Q5: Fix Loop Capture
========================================
Expected Output:
0 1 2 (any order)
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
Q8: Buffered Channel Trap
========================================
Expected Output:
(no block)
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
1 2
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
