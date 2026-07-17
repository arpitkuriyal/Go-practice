package main

import "fmt"

func reverse(str string) string {

	// can't do this because string in go is immutable
	// for i := 0; i < len(str); i++ {
	// 	str[i], str[len(str)-i] = str[len(str)-i], str[i]
	// }
	// return str

	b := []byte(str)
	left := 0
	right := len(b) - 1
	for left < right {
		b[right], b[left] = b[left], b[right]
		left++
		right--
	}
	return string(b)
}

func main() {
	var s string
	fmt.Scan(&s)

	ans := reverse(s)
	fmt.Println(ans)
}
