// Package shapecalculator demonstrates interface-based polymorphism.
package shapecalculator

import (
	"errors"
	"math"
)

var ErrInvalidDimensions = errors.New("shape dimensions must be positive and valid")

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{ Radius float64 }

func NewCircle(radius float64) (Circle, error) {
	if radius <= 0 {
		return Circle{}, ErrInvalidDimensions
	}
	return Circle{Radius: radius}, nil
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rectangle struct{ Width, Height float64 }

func NewRectangle(width, height float64) (Rectangle, error) {
	if width <= 0 || height <= 0 {
		return Rectangle{}, ErrInvalidDimensions
	}
	return Rectangle{Width: width, Height: height}, nil
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Triangle struct{ A, B, C float64 }

func NewTriangle(a, b, c float64) (Triangle, error) {
	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || a+c <= b || b+c <= a {
		return Triangle{}, ErrInvalidDimensions
	}
	return Triangle{A: a, B: b, C: c}, nil
}

func (t Triangle) Area() float64 {
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}
func (t Triangle) Perimeter() float64 { return t.A + t.B + t.C }

func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}
