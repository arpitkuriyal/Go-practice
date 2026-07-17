// Package sliceoperations contains non-mutating, generic slice helpers.
package sliceoperations

import "errors"

var ErrIndexOutOfRange = errors.New("index out of range")

// Insert returns a new slice with value placed at index.
func Insert[T any](items []T, index int, value T) ([]T, error) {
	if index < 0 || index > len(items) {
		return nil, ErrIndexOutOfRange
	}

	result := make([]T, 0, len(items)+1)
	result = append(result, items[:index]...)
	result = append(result, value)
	result = append(result, items[index:]...)
	return result, nil
}

// Remove returns a new slice without the item at index.
func Remove[T any](items []T, index int) ([]T, error) {
	if index < 0 || index >= len(items) {
		return nil, ErrIndexOutOfRange
	}

	result := make([]T, 0, len(items)-1)
	result = append(result, items[:index]...)
	result = append(result, items[index+1:]...)
	return result, nil
}

// Filter returns a new slice containing values accepted by keep.
func Filter[T any](items []T, keep func(T) bool) []T {
	result := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			result = append(result, item)
		}
	}
	return result
}

// Unique preserves the first occurrence of every comparable value.
func Unique[T comparable](items []T) []T {
	seen := make(map[T]struct{}, len(items))
	result := make([]T, 0, len(items))
	for _, item := range items {
		if _, exists := seen[item]; exists {
			continue
		}
		seen[item] = struct{}{}
		result = append(result, item)
	}
	return result
}
