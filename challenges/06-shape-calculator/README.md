# Shape Calculator: Interfaces and Polymorphism

## Challenge

Create circles, rectangles, and triangles that expose a common `Shape` behavior. Validate dimensions and calculate the total area of mixed shapes.

## Core interface

```go
type Shape interface {
	Area() float64
	Perimeter() float64
}
```

Each concrete type satisfies this interface implicitly by implementing both methods. No `implements` keyword is required.

```go
func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}
```

## Concepts practised

- Small behavior-focused interfaces.
- Value receivers for small immutable shape values.
- Constructors that enforce valid dimensions.
- Heron’s formula for a triangle’s area.
- Polymorphism: `TotalArea` does not need to know the concrete shape type.

## Validation rules

- Circle radius, rectangle width, and height must be positive.
- Triangle sides must be positive and satisfy the triangle inequality: each pair must be greater than the remaining side.
- Callers can use `errors.Is(err, ErrInvalidDimensions)` for a rejected constructor.

## Complexity

Area/perimeter calculations are `O(1)`. `TotalArea` is `O(n)` for `n` shapes.

## Interview answer

“I define an interface for the behavior the caller needs, not for a concrete type hierarchy. Each shape implements it implicitly, so code that sums areas is open to new shapes without modification.”

## Test

```bash
go test ./challenges/06-shape-calculator
```
