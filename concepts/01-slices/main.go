package main

import "fmt"

/*
========================================
Q1: Basic Value vs Pointer
========================================
Expected Output:
10
100
*/
func changeVal(x int) {
	x = 100
}

func changePtr(x *int) {
	*x = 100
}

/*
========================================
Q2: Slice Modification (Underlying Array)
========================================
Expected Output:
[100 2 3]
*/
func modifySlice(s []int) {
	s[0] = 100
}

/*
========================================
Q3: Append WITHOUT Returning (Trap)
========================================
Expected Output:
[1 2 3]
*/
func appendNoReturn(s []int) {
	s = append(s, 100)
}

/*
========================================
Q4: Append WITH Return
========================================
Expected Output:
[1 2 3 100]
*/
func appendWithReturn(s []int) []int {
	return append(s, 100)
}

/*
========================================
Q5: Capacity Trap (No New Array)
========================================
Expected Output:
[1 2 3 100 0]
*/
func capNoRealloc() {
	arr := make([]int, 3, 5)
	arr[0], arr[1], arr[2] = 1, 2, 3

	s := arr
	s = append(s, 100)

	fmt.Println(arr)
}

/*
========================================
Q6: Capacity Full (New Array Created)
========================================
Expected Output:
[1 2 3]
[1 2 3 100]
*/
func capRealloc() {
	arr := []int{1, 2, 3}

	s := arr
	s = append(s, 100)

	fmt.Println(arr)
	fmt.Println(s)
}

/*
========================================
Q7: Slice Length NOT Updated Outside
========================================
Expected Output:
inside: [1 2 3 100]
outside: [1 2 3]
*/
func lengthTrap(s []int) {
	s = append(s, 100)
	fmt.Println("inside:", s)
}

/*
========================================
Q8: Shared Underlying Array Trap
========================================
Expected Output:
[1 2]
[100 2]
*/
func sharedArray() {
	arr := []int{1, 2, 3}

	a := arr[:2]
	b := arr[:2]

	a[0] = 100

	fmt.Println(a)
	fmt.Println(b)
}

/*
========================================
Q9: Append Overwrites Other Slice (Danger)
========================================
Expected Output:
[1 2 100]
[1 2]
*/
func overwriteTrap() {
	arr := make([]int, 2, 4)
	arr[0], arr[1] = 1, 2

	a := arr
	b := arr

	a = append(a, 100)

	fmt.Println(a)
	fmt.Println(b)
}

/*
========================================
Q10: Struct Passed by Value vs Pointer
========================================
Expected Output:
{Alice}
{Bob}
*/
type User struct {
	name string
}

func changeUser(u User) {
	u.name = "Bob"
}

func changeUserPtr(u *User) {
	u.name = "Bob"
}

func main() {

	// Q1
	a := 10
	changeVal(a)
	fmt.Println(a) // 10

	changePtr(&a)
	fmt.Println(a) // 100

	// Q2
	s := []int{1, 2, 3}
	modifySlice(s)
	fmt.Println(s) // [100 2 3]

	// Q3
	s2 := []int{1, 2, 3}
	appendNoReturn(s2)
	fmt.Println(s2) // [1 2 3]

	// Q4
	s3 := []int{1, 2, 3}
	s3 = appendWithReturn(s3)
	fmt.Println(s3) // [1 2 3 100]

	// Q5
	capNoRealloc()

	// Q6
	capRealloc()

	// Q7
	x := []int{1, 2, 3}
	lengthTrap(x)
	fmt.Println("outside:", x)

	// Q8
	sharedArray()

	// Q9
	overwriteTrap()

	// Q10
	u := User{name: "Alice"}
	changeUser(u)
	fmt.Println(u) // {Alice}

	changeUserPtr(&u)
	fmt.Println(u) // {Bob}
}
