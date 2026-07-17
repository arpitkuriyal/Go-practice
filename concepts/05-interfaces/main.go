package main

import "fmt"

/*
========================================
Q1: Nil Interface Trap
========================================
Expected Output:
false
*/
func q1() {
	var p *int = nil
	var i interface{} = p

	fmt.Println(i == nil)
}

/*
========================================
Q2: Function Nil Check Bug
========================================
Expected Output:
not nil
*/
func check(i interface{}) {
	if i == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

func q2() {
	var p *int = nil
	check(p)
}

/*
========================================
Q3: Interface Comparison
========================================
Expected Output:
true
*/
func q3() {
	var a interface{} = 10
	var b interface{} = 10

	fmt.Println(a == b)
}

/*
========================================
Q4: Slice Inside Interface
========================================
Expected Output:
panic
*/
func q4() {
	var a interface{} = []int{1}
	var b interface{} = []int{1}

	fmt.Println(a == b)
}

/*
========================================
Q5: Type Assertion Panic
========================================
Expected Output:
panic
*/
func q5() {
	var i interface{} = 10

	x := i.(string)
	fmt.Println(x)
}

/*
========================================
Q6: Safe Type Assertion
========================================
Expected Output:
"" false
*/
func q6() {
	var i interface{} = 10

	x, ok := i.(string)
	fmt.Println(x, ok)
}

/*
========================================
Q7: Pointer Receiver Trap
========================================
Expected:
compile error
*/
type Speaker interface {
	Speak()
}

type Dog struct{}

func (d *Dog) Speak() {}

func q7() {
	// var s Speaker = Dog{} // ❌ compile error
	// fmt.Println(s)
}

/*
========================================
Q8: Nil Pointer Inside Interface
========================================
Expected Output:
false
panic
*/
type User struct{}

func (u *User) Name() string {
	return "Arpit"
}

func q8() {
	var u *User = nil
	var i interface{} = u

	fmt.Println(i == nil)
	fmt.Println(i.(*User).Name())
}

/*
========================================
Q9: Type Switch
========================================
Expected Output:
int
*/
func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

func q9() {
	checkType(10)
}

/*
========================================
Q10: Value vs Pointer Implementation
========================================
Expected Output:
valid
*/
type T struct{}

func (t T) Speak() {}

type Speaker2 interface {
	Speak()
}

func q10() {
	var s Speaker2 = &T{}
	fmt.Println(s)
}

func main() {
	q1()
	q2()
	q3()
	// q4() // will panic
	// q5() // will panic
	q6()
	// q7() // compile error
	// q8() // will panic
	q9()
	q10()
}
