package main

import "fmt"

/*
========================================
Q1: Nil Map Assignment
========================================
Expected:
panic
*/
func q1() {
	var m map[string]int
	m["a"] = 1
}

/*
========================================
Q2: Nil Map Read
========================================
Expected Output:
0
*/
func q2() {
	var m map[string]int
	fmt.Println(m["a"])
}

/*
========================================
Q3: Map Initialization
========================================
Expected Output:
map[a:1]
*/
func q3() {
	m := make(map[string]int)
	m["a"] = 1
	fmt.Println(m)
}

/*
========================================
Q4: Map Passed to Function
========================================
Expected Output:
map[a:100]
*/
func modify(m map[string]int) {
	m["a"] = 100
}

func q4() {
	m := map[string]int{"a": 1}
	modify(m)
	fmt.Println(m)
}

/*
========================================
Q5: Map Reassignment Trap
========================================
Expected Output:
map[a:1]
*/
func reassign(m map[string]int) {
	m = make(map[string]int)
	m["a"] = 100
}

func q5() {
	m := map[string]int{"a": 1}
	reassign(m)
	fmt.Println(m)
}

/*
========================================
Q6: Key Existence Check
========================================
Expected Output:
0 false
*/
func q6() {
	m := map[string]int{}

	val, ok := m["a"]
	fmt.Println(val, ok)
}

/*
========================================
Q7: Iteration Order
========================================
Expected:
random order
*/
func q7() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
========================================
Q8: Delete Key
========================================
Expected Output:
map[]
*/
func q8() {
	m := map[string]int{"a": 1}
	delete(m, "a")
	fmt.Println(m)
}

/*
========================================
Q9: Map of Slices Trap
========================================
Expected:
panic
*/
func q9() {
	m := make(map[string][]int)
	m["a"][0] = 1
}

/*
========================================
Q10: Fix Map of Slices
========================================
Expected Output:
map[a:[1]]
*/
func q10() {
	m := make(map[string][]int)
	m["a"] = append(m["a"], 1)
	fmt.Println(m)
}

func main() {
	// q1() // panic
	q2()
	q3()
	q4()
	q5()
	q6()
	q7()
	q8()
	// q9() // panic
	q10()
}
