package main

import "fmt"

// sum of numbers -> first simple problem
func main() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b, &c)

	if err != nil {
		fmt.Println("invalid input. Must be number only")
		return
	}

	ans1 := sum2(a, b)
	fmt.Println("sum of 2 is :- ", ans1)

	ans2 := sum3(a, b, c)
	fmt.Println("sum of 3 is :-", ans2)

	ans := sum(1, 2, 3, 4, 5)
	fmt.Println("sum is := ", ans)
}

func sum2(a, b int) int {
	return a + b
}

func sum3(a, b, c int) int {
	return a + b + c
}

func sum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
