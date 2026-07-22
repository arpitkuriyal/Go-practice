package main

import "fmt"

/*
==================================================
Q1: Why is this FALSE even though the pointer is nil?

Question:
What will this print?

Rule:
An interface is nil ONLY when:
1. Dynamic type == nil
2. Dynamic value == nil

Here:
type  = *int
value = nil

So interface != nil.
==================================================
Expected Output:
false
*/
func q1() {
	var p *int = nil
	var i interface{} = p

	fmt.Println(i == nil)
}

/*
==================================================
Q2: Why does this print "not nil"?

Question:
Why doesn't check() detect nil?

Rule:
Passing a nil pointer into interface{} creates:

(type=*int, value=nil)

which is NOT a nil interface.
==================================================
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
==================================================
Q3: Can two interface values be compared?

Question:
What will this print?

Rule:
Two interfaces are comparable when their underlying
dynamic values are comparable.

int is comparable.

(type=int,value=10) == (type=int,value=10)
==================================================
Expected Output:
true
*/
func q3() {
	var a interface{} = 10
	var b interface{} = 10

	fmt.Println(a == b)
}

/*
==================================================
Q4: Why does this panic?

Question:
Can interface values holding slices be compared?

Rule:
Interfaces are comparable ONLY if their underlying
dynamic values are comparable.

Slices are NOT comparable.

Comparing them causes a panic.
==================================================
Expected Output:
panic:
comparing uncomparable type []int
*/
func q4() {
	var a interface{} = []int{1}
	var b interface{} = []int{1}

	fmt.Println(a == b)
}

/*
==================================================
Q5: Type Assertion Panic

Question:
What happens here?

Rule:
A single-value type assertion MUST succeed.

If the stored type doesn't match,
Go panics.

Stored type : int
Requested   : string
==================================================
Expected Output:
panic:
interface conversion: interface {} is int, not string
*/
func q5() {
	var i interface{} = 10

	x := i.(string)
	fmt.Println(x)
}

/*
==================================================
Q6: Safe Type Assertion

Question:
How do we avoid the panic?

Rule:
Use the comma-ok idiom.

value, ok := x.(T)

If assertion fails:
value = zero value
ok = false
==================================================
Expected Output:

	false
*/
func q6() {
	var i interface{} = 10

	x, ok := i.(string)
	fmt.Println(x, ok)
}

/*
==================================================
Q7: Pointer Receiver Trap

Question:
Why doesn't this compile?

Rule:
Method set:

Value methods      -> available on T and *T
Pointer methods    -> available ONLY on *T

Dog has only a pointer receiver,
so Dog{} does NOT implement Speaker.

*Dog DOES implement Speaker.
==================================================
Expected:
compile error
*/
type Speaker interface {
	Speak()
}

type Dog struct{}

func (d *Dog) Speak() {}

func q7() {
	// var s Speaker = Dog{} // compile error
	// var s Speaker = &Dog{} // correct
}

/*
==================================================
Q8: Nil Pointer Inside Interface

Question:
Why is the interface not nil?
Why does calling Name() panic?

Rule:
The interface contains:

type=*User
value=nil

So interface != nil.

Method call dereferences a nil pointer.
==================================================
Expected Output:
false
panic:
nil pointer dereference
*/
type User struct {
	name string
}

func (u *User) Name() string {
	return u.name
}

func q8() {
	var u *User = nil
	var i interface{} = u

	fmt.Println(i == nil)
	fmt.Println(i.(*User).Name())
}

/*
==================================================
Q9: Type Switch

Question:
How do we check an interface's dynamic type?

Rule:
Use a type switch.

switch x.(type)

No explicit type assertion needed.
==================================================
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
==================================================
Q10: Value vs Pointer Method Set

Question:
Why does this work?

Rule:
Value receiver methods belong to BOTH:

T
*T

So both implement the interface.

If Speak() had a pointer receiver,
only *T would implement Speaker2.
==================================================
Expected Output:
{}
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
	fmt.Println("===== Q1: Nil Interface =====")
	q1()

	fmt.Println("\n===== Q2: Nil Check Bug =====")
	q2()

	fmt.Println("\n===== Q3: Comparable Interface =====")
	q3()

	fmt.Println("\n===== Q4: Slice Comparison =====")
	// q4() // panic

	fmt.Println("\n===== Q5: Unsafe Type Assertion =====")
	// q5() // panic

	fmt.Println("\n===== Q6: Safe Type Assertion =====")
	q6()

	fmt.Println("\n===== Q7: Pointer Receiver =====")
	// q7() // compile error

	fmt.Println("\n===== Q8: Nil Pointer in Interface =====")
	// q8() // panic

	fmt.Println("\n===== Q9: Type Switch =====")
	q9()

	fmt.Println("\n===== Q10: Method Set =====")
	q10()
}
