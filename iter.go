package slices

import "iter"

func Iter[T any](sl []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range sl {
			if !yield(i, v) {
				return
			}
		}
	}
}
