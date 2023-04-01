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

// Find returns the index of the value in sl or -1 if not found.
func Find[T comparable](sl []T, v T) int {
	return FindFromFunc(sl, v, 0, cmpFn[T])
}

// FindFunc returns the index of the value in sl or -1 if not found using
// a function to compare values.
func FindFunc[T any](sl []T, v T, f CmpFunc[T]) int {
	return FindFromFunc(sl, v, 0, f)
}

// FindFrom returns the index of the value in sl or -1 if not found.
// It starts at index start.
func FindFrom[T comparable](sl []T, v T, start int) int {
	return FindFromFunc(sl, v, start, cmpFn[T])
}

// FindFromFunc returns the index of the value in sl or -1 if not found using
// a function to compare values. It starts at index start.
func FindFromFunc[T any](sl []T, v T, start int, f CmpFunc[T]) int {
	for i := start; i < len(sl); i++ {
		if f(v, sl[i]) {
			return i + start
		}
	}
	return -1
}

// FindAll returns all indexes of the value v.
func FindAll[T comparable](sl []T, v T) []int {
	return FindAllFunc(sl, v, cmpFn[T])
}

// FindAll returns all indexes of the value v using a function to compare values.
func FindAllFunc[T any](sl []T, v T, f CmpFunc[T]) []int {
	if sl == nil {
		return nil
	}
	result := []int{}
	for i, x := range sl {
		if f(v, x) {
			result = append(result, i)
		}
	}
	return result
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

// TODO
// func Group[T comparable](sl []T) [][]T

// func GroupFunc[T any](sl []T, f CmpFunc[T]) [][]T

// func Purge[T comparable](sl, ps []T) []T

// func PurgeFunc[T any](sl, ps []T, f CmpFunc[T]) []T
