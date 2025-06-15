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
