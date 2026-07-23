package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// Birthday has a pointer receiver because it changes the original user.
func (u *User) Birthday() {
	u.Age++
}

// Greeting has a value receiver because it only reads the user.
func (u User) Greeting() string {
	return "Hello, " + u.Name
}

type Employee struct {
	User // embedded field: Employee gets User's fields and methods
	Role string
}

func main() {
	user := User{Name: "Arpit", Age: 23}
	fmt.Println(user.Greeting())

	user.Birthday()
	fmt.Println("age after birthday:", user.Age)

	employee := Employee{
		User: User{Name: "Neha", Age: 25},
		Role: "Engineer",
	}
	fmt.Println(employee.Name, "is an", employee.Role)
}
