package sliceoperations

import (
	"errors"
	"reflect"
	"testing"
)

func TestInsertAndRemove(t *testing.T) {
	original := []int{1, 3, 4}
	inserted, err := Insert(original, 1, 2)
	if err != nil || !reflect.DeepEqual(inserted, []int{1, 2, 3, 4}) {
		t.Fatalf("Insert() = %v, %v", inserted, err)
	}
	if !reflect.DeepEqual(original, []int{1, 3, 4}) {
		t.Fatalf("Insert mutated input: %v", original)
	}

	removed, err := Remove(inserted, 2)
	if err != nil || !reflect.DeepEqual(removed, []int{1, 2, 4}) {
		t.Fatalf("Remove() = %v, %v", removed, err)
	}
}

func TestSliceHelpers(t *testing.T) {
	if _, err := Remove([]int{}, 0); !errors.Is(err, ErrIndexOutOfRange) {
		t.Fatalf("Remove empty error = %v", err)
	}
	if got := Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 }); !reflect.DeepEqual(got, []int{2, 4}) {
		t.Fatalf("Filter() = %v", got)
	}
	if got := Unique([]string{"go", "go", "rust", "go"}); !reflect.DeepEqual(got, []string{"go", "rust"}) {
		t.Fatalf("Unique() = %v", got)
	}
}
