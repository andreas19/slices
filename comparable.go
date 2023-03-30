package slices

// A CmpFunc compares two values and returns true if they are equal.
type CmpFunc[T any] func(v1, v2 T) bool

func cmpFn[T comparable](v1, v2 T) bool {
	return v1 == v2
}

// Contains returns true if the value is in the slice.
func Contains[T comparable](sl []T, v T) bool {
	return ContainsFunc(sl, v, cmpFn[T])
}

// ContainsFunc returns true if the value is in the slice using a function to compare values.
func ContainsFunc[T any](sl []T, v T, f CmpFunc[T]) bool {
	return FindFunc(sl, v, f) != -1
}

// Equal returns true if the slices are equal.
func Equal[T comparable](sl1, sl2 []T) bool {
	return EqualFunc(sl1, sl2, cmpFn[T])
}

// EqualFunc returns true if the slices are equal using a function to compare values.
func EqualFunc[T any](sl1, sl2 []T, f CmpFunc[T]) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i := range sl1 {
		if !f(sl1[i], sl2[i]) {
			return false
		}
	}
	return true
}

// Find returns the index of the value or -1 if not found.
func Find[T comparable](sl []T, v T) int {
	return FindFunc(sl, v, cmpFn[T])
}

// FindFunc returns the index of the value or -1 if not found using a function to compare values.
func FindFunc[T any](sl []T, v T, f CmpFunc[T]) int {
	for i, x := range sl {
		if f(x, v) {
			return i
		}
	}
	return -1
}

// Unique returns the unique values from a slice.
func Unique[T comparable](sl []T) []T {
	if sl == nil {
		return nil
	}
	result := []T{}
	seen := make(map[T]struct{})
	for _, v := range sl {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// UniqueFunc returns the unique values from a slice using a function to compare values.
func UniqueFunc[T any](sl []T, f CmpFunc[T]) []T {
	if sl == nil {
		return nil
	}
	result := []T{}
	for _, v := range sl {
		if !ContainsFunc(result, v, f) {
			result = append(result, v)
		}
	}
	return result
}
