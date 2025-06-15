package slices

import "iter"

// Iter returns an iterator over index-values pairs in the slice.
func Iter[T any](sl []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range sl {
			if !yield(i, v) {
				return
			}
		}
	}
}

// Values returns an iterator over the values in the slice.
func Values[T any](sl []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range sl {
			if !yield(v) {
				return
			}
		}
	}
}

// Collect returns a new slice with the values from seq.
func Collect[T any](seq iter.Seq[T]) []T {
	result := []T{}
	for v := range seq {
		result = append(result, v)
	}
	return result
}
