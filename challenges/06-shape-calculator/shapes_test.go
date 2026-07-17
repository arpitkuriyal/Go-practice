package shapecalculator

import (
	"errors"
	"math"
	"testing"
)

func TestShapes(t *testing.T) {
	circle, _ := NewCircle(1)
	rectangle, _ := NewRectangle(3, 4)
	triangle, _ := NewTriangle(3, 4, 5)

	if got, want := TotalArea([]Shape{circle, rectangle, triangle}), math.Pi+18; math.Abs(got-want) > 1e-9 {
		t.Fatalf("TotalArea() = %v, want %v", got, want)
	}
	if got, want := rectangle.Perimeter(), 14.0; got != want {
		t.Fatalf("Perimeter() = %v, want %v", got, want)
	}
}

func TestInvalidDimensions(t *testing.T) {
	if _, err := NewCircle(0); !errors.Is(err, ErrInvalidDimensions) {
		t.Fatalf("NewCircle(0) error = %v", err)
	}
	if _, err := NewTriangle(1, 2, 4); !errors.Is(err, ErrInvalidDimensions) {
		t.Fatalf("NewTriangle invalid sides error = %v", err)
	}
}
