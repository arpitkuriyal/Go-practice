package main

import "fmt"

func main() {
	q1()
	q2()
	q3()
	q4()
	q5()
	q6()
	q7()
	q8()
	q9()
	q10()
}

/////////////////////////////////////////////////////////
// Q1: Value vs Pointer
/////////////////////////////////////////////////////////

func changeVal(x int) {
	x = 100
}

func changePtr(x *int) {
	*x = 100
}

func q1() {
	fmt.Println("===== Q1: Value vs Pointer =====")

	a := 10

	changeVal(a)
	fmt.Println(a)

	changePtr(&a)
	fmt.Println(a)

	// Expected:
	// 10
	// 100
	//
	// Lesson:
	// Values are copied.
	// Pointers modify the original variable.
}

/////////////////////////////////////////////////////////
// Q2: Slice shares backing array
/////////////////////////////////////////////////////////

func modifySlice(s []int) {
	s[0] = 100
}

func q2() {
	fmt.Println("\n===== Q2: Slice Modification =====")

	s := []int{1, 2, 3}

	modifySlice(s)

	fmt.Println(s)

	// Expected:
	// [100 2 3]
	//
	// Lesson:
	// Slice header is copied,
	// backing array is shared.
}

/////////////////////////////////////////////////////////
// Q3: append WITHOUT returning
/////////////////////////////////////////////////////////

func appendNoReturn(s []int) {
	s = append(s, 100)
}

func q3() {
	fmt.Println("\n===== Q3: append WITHOUT return =====")

	s := []int{1, 2, 3}

	appendNoReturn(s)

	fmt.Println(s)

	// Expected:
	// [1 2 3]
	//
	// Lesson:
	// append returns a NEW slice header.
	// Ignoring it means the caller is unchanged.
}

/////////////////////////////////////////////////////////
// Q4: append WITH return
/////////////////////////////////////////////////////////

func appendWithReturn(s []int) []int {
	return append(s, 100)
}

func q4() {
	fmt.Println("\n===== Q4: append WITH return =====")

	s := []int{1, 2, 3}

	s = appendWithReturn(s)

	fmt.Println(s)

	// Expected:
	// [1 2 3 100]
	//
	// Lesson:
	// Always keep the slice returned by append.
}

/////////////////////////////////////////////////////////
// Q5: append reuses the backing array (capacity available)
/////////////////////////////////////////////////////////

func q5() {
	fmt.Println("\n===== Q5: Existing Backing Array =====")

	arr := make([]int, 3, 5)
	arr[0], arr[1], arr[2] = 1, 2, 3

	s := arr
	s = append(s, 100)

	fmt.Println("arr      :", arr)
	fmt.Println("arr full :", arr[:cap(arr)])
	fmt.Println("s        :", s)

	// Output:
	// arr      : [1 2 3]
	// arr full : [1 2 3 100 0]
	// s        : [1 2 3 100]
	//
	// Lesson:
	// append reused the same backing array because
	// there was spare capacity.
	//
	// arr still has len = 3,
	// while s has len = 4.
}

/////////////////////////////////////////////////////////
// Q6: append allocates a NEW backing array
/////////////////////////////////////////////////////////

func q6() {
	fmt.Println("\n===== Q6: New Backing Array =====")

	arr := []int{1, 2, 3}

	s := arr
	s = append(s, 100)

	fmt.Println("arr :", arr)
	fmt.Println("s   :", s)

	// Output:
	// arr : [1 2 3]
	// s   : [1 2 3 100]
	//
	// Lesson:
	// Capacity was full.
	// append allocated a new backing array.
	//
	// arr and s no longer share memory.
}

/////////////////////////////////////////////////////////
// Q7: Slice header is copied
/////////////////////////////////////////////////////////

func lengthTrap(s []int) {
	s = append(s, 100)

	fmt.Println("inside :", s)
}

func q7() {
	fmt.Println("\n===== Q7: Slice Header Is Copied =====")

	x := make([]int, 3, 5)
	x[0], x[1], x[2] = 1, 2, 3

	lengthTrap(x)

	fmt.Println("outside:", x)
	fmt.Println("full   :", x[:cap(x)])

	// Output:
	// inside : [1 2 3 100]
	// outside: [1 2 3]
	// full   : [1 2 3 100 0]
	//
	// Lesson:
	// The backing array is shared,
	// but the slice header (len, cap)
	// is copied into the function.
}

/////////////////////////////////////////////////////////
// Q8: Two slices share one backing array
/////////////////////////////////////////////////////////

func q8() {
	fmt.Println("\n===== Q8: Shared Backing Array =====")

	arr := []int{1, 2, 3}

	a := arr[:2]
	b := arr[:2]

	a[0] = 100

	fmt.Println("a :", a)
	fmt.Println("b :", b)
	fmt.Println("arr:", arr)

	// Output:
	// a : [100 2]
	// b : [100 2]
	// arr: [100 2 3]
	//
	// Lesson:
	// a and b point to the same backing array.
	// Modifying one affects the other.
}

/////////////////////////////////////////////////////////
// Q9: append can affect another slice
/////////////////////////////////////////////////////////

func q9() {
	fmt.Println("\n===== Q9: append Can Affect Another Slice =====")

	arr := make([]int, 2, 4)
	arr[0], arr[1] = 1, 2

	a := arr
	b := arr

	a = append(a, 100)

	fmt.Println("a      :", a)
	fmt.Println("b      :", b)
	fmt.Println("b full :", b[:cap(b)])

	// Output:
	// a      : [1 2 100]
	// b      : [1 2]
	// b full : [1 2 100 0]
	//
	// Lesson:
	// append reused the backing array.
	//
	// b still has len = 2,
	// but the backing array now contains 100.
}

/////////////////////////////////////////////////////////
// Q10: Struct - Value vs Pointer
/////////////////////////////////////////////////////////

type User struct {
	name string
}

func changeUser(u User) {
	u.name = "Bob"
}

func changeUserPtr(u *User) {
	u.name = "Bob"
}

func q10() {
	fmt.Println("\n===== Q10: Struct Value vs Pointer =====")

	u := User{name: "Alice"}

	changeUser(u)
	fmt.Println("after value  :", u)

	changeUserPtr(&u)
	fmt.Println("after pointer:", u)

	// Output:
	// after value  : {Alice}
	// after pointer: {Bob}
	//
	// Lesson:
	// Passing a struct copies the entire struct.
	// Changes affect only the copy.
	//
	// Passing a pointer allows the function
	// to modify the original struct.
}
