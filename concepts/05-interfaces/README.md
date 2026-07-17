# 🚨 Go Interfaces — Trap Questions + Follow-ups

---

# 🔥 SECTION 1: Trap Questions (Try Before Seeing Answers)

---

## ❓ Q1: Nil Interface Trap

```go
var p *int = nil
var i interface{} = p

fmt.Println(i == nil)
```

👉 What is the output?

---

## ❓ Q2: Function Nil Check Bug

```go
func check(i interface{}) {
	if i == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

var p *int = nil
check(p)
```

👉 What prints?

---

## ❓ Q3: Interface Comparison

```go
var a interface{} = 10
var b interface{} = 10

fmt.Println(a == b)
```

👉 Output?

---

## ❓ Q4: Slice Inside Interface

```go
var a interface{} = []int{1}
var b interface{} = []int{1}

fmt.Println(a == b)
```

👉 What happens?

---

## ❓ Q5: Type Assertion Panic

```go
var i interface{} = 10

x := i.(string)
fmt.Println(x)
```

👉 Output?

---

## ❓ Q6: Safe Type Assertion

```go
var i interface{} = 10

x, ok := i.(string)
fmt.Println(x, ok)
```

👉 Output?

---

## ❓ Q7: Pointer Receiver Trap

```go
type Speaker interface {
	Speak()
}

type Dog struct{}

func (d *Dog) Speak() {}

var s Speaker = Dog{}
```

👉 Compile or error?

---

## ❓ Q8: Nil Pointer Inside Interface Method Call

```go
type User struct{}

func (u *User) Name() string {
	return "Arpit"
}

var u *User = nil
var i interface{} = u

fmt.Println(i == nil)
fmt.Println(i.(*User).Name())
```

👉 What happens?

---

## ❓ Q9: Type Switch

```go
func check(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

check(10)
```

👉 Output?

---

## ❓ Q10: Value vs Pointer Implementation

```go
type T struct{}

func (t T) Speak() {}

type Speaker interface {
	Speak()
}

var s Speaker = &T{}
```

👉 Valid or not?

---

---

# ✅ SECTION 2: Follow-ups & Explanations

---

## 🔹 Q1 Follow-up: Why `i != nil`?

👉 Interface = `(type, value)`

* Here: `(*int, nil)`
* Since type exists → interface is NOT nil

---

## 🔹 Q2 Follow-up: Real Bug

👉 Output: `not nil`

* Function sees `(type=*int, value=nil)`
* Common production bug in APIs

---

## 🔹 Q3 Follow-up: Comparable Types

👉 Output: `true`

* Both have same type (`int`) and value (`10`)

---

## 🔹 Q4 Follow-up: Why Panic?

👉 Reason:

* Slice is NOT comparable
* Interface tries to compare underlying types → runtime panic

---

## 🔹 Q5 Follow-up: Type Assertion Panic

👉 Panic:

* `int` cannot be asserted as `string`

---

## 🔹 Q6 Follow-up: Safe Assertion

👉 Output:

```text
"" false
```

* No panic
* `ok` tells success

---

## 🔹 Q7 Follow-up: Pointer Receiver Rule

👉 ❌ Compile Error

Rule:

* Method defined on `*Dog`
* Only `*Dog` implements interface

✅ Fix:

```go
var s Speaker = &Dog{}
```

---

## 🔹 Q8 Follow-up: Hidden Panic

👉 Output:

```text
false
panic
```

Why?

* Interface ≠ nil
* But underlying pointer is nil → method call crashes

---

## 🔹 Q9 Follow-up: Type Switch

👉 Output:

```text
int
```

* Matches concrete type

---

## 🔹 Q10 Follow-up: Method Set Rule

👉 ✅ Valid

Rule:

* Value receiver (`T`) works for both:

  * `T`
  * `*T`

---

# 🧠 Final Mental Model

* Interface = `(type, value)`
* Nil check = both must be nil
* Comparison depends on underlying type
* Pointer receiver ≠ value receiver
* Type assertion can panic

---

# 🚀 Interview Strategy

Whenever you see interface:

1. What is the underlying type?
2. Is it truly nil?
3. Is this type comparable?
4. Pointer or value receiver?

👉 These 4 questions solve almost every interface problem.
