package patternmatching

import (
	"reflect"
	"testing"
)

func TestFindAll(t *testing.T) {
	tests := []struct {
		text, pattern string
		want          []int
	}{
		{"banana", "ana", []int{1, 3}},
		{"hello", "z", []int{}},
		{"hello", "", []int{}},
	}
	for _, test := range tests {
		if got := FindAll(test.text, test.pattern); !reflect.DeepEqual(got, test.want) {
			t.Errorf("FindAll(%q, %q) = %v, want %v", test.text, test.pattern, got, test.want)
		}
	}
}

func TestFindAnagramIndices(t *testing.T) {
	if got, want := FindAnagramIndices("cbaebabacd", "abc"), []int{0, 6}; !reflect.DeepEqual(got, want) {
		t.Fatalf("FindAnagramIndices() = %v, want %v", got, want)
	}
	if got, want := FindAnagramIndices("éaé", "éa"), []int{0, 1}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Unicode FindAnagramIndices() = %v, want %v", got, want)
	}
}
