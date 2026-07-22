package testingtooling

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{"empty", nil, 0},
		{"positive values", []int{1, 2, 3}, 6},
		{"negative values", []int{-2, 5}, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Sum(test.values...); got != test.want {
				t.Fatalf("Sum(%v) = %d, want %d", test.values, got, test.want)
			}
		})
	}
}

type fakeStore struct {
	exists bool
	err    error
	got    string
}

func (s *fakeStore) Exists(_ context.Context, email string) (bool, error) {
	s.got = email
	return s.exists, s.err
}

func TestRegistrationServiceUsesFake(t *testing.T) {
	store := &fakeStore{}
	service := NewRegistrationService(store)
	if err := service.Register(context.Background(), " Arpit@Example.COM "); err != nil {
		t.Fatal(err)
	}
	if store.got != "arpit@example.com" {
		t.Fatalf("store received %q", store.got)
	}

	store.exists = true
	if err := service.Register(context.Background(), "arpit@example.com"); !errors.Is(err, ErrEmailTaken) {
		t.Fatalf("Register() error = %v, want ErrEmailTaken", err)
	}
}

func BenchmarkSum(b *testing.B) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ReportAllocs()
	for b.Loop() {
		_ = Sum(values...)
	}
}

func FuzzNormalizeEmail(f *testing.F) {
	f.Add(" Arpit@Example.COM ")
	f.Add("")
	f.Fuzz(func(t *testing.T, input string) {
		got := NormalizeEmail(input)
		if got != strings.ToLower(strings.TrimSpace(got)) {
			t.Fatalf("NormalizeEmail(%q) is not normalized: %q", input, got)
		}
	})
}
