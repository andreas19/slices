package slices

import "sort"

type Ordered interface {
	Number | ~string
}

// A LessFunc compares two values and returns true if the first is less then the second.
type LessFunc[T any] func(v1, v2 T) bool

func lessFn[T Ordered](v1, v2 T) bool {
	return v1 < v2
}

// Min returns the minimum from the slice.
// Panics if the slice is empty or nil.
func Min[T Ordered](sl []T) T {
	return MinFunc(sl, lessFn[T])
}

// MinFunc returns the minimum from the slice using a function to compare values.
// Panics if the slice is empty or nil.
func MinFunc[T any](sl []T, f LessFunc[T]) T {
	if len(sl) == 0 {
		panic("slice is empty or nil")
	}
	result := sl[0]
	for _, v := range sl[1:] {
		if f(v, result) {
			result = v
		}
	}
	return result
}

// Max returns the maximum from the slice.
// Panics if the slice is empty or nil.
func Max[T Ordered](sl []T) T {
	return MaxFunc(sl, lessFn[T])
}

// MaxFunc returns the maximum from the slice using a function to compare values.
// Panics if the slice is empty or nil.
func MaxFunc[T any](sl []T, f LessFunc[T]) T {
	if len(sl) == 0 {
		panic("slice is empty or nil")
	}
	result := sl[0]
	for _, v := range sl[1:] {
		if f(result, v) {
			result = v
		}
	}
	return result
}

// Extrema returns minimum and maximum from the slice.
// Panics if the slice is empty or nil.
func Extrema[T Ordered](sl []T) (min, max T) {
	return ExtremaFunc(sl, lessFn[T])
}

// ExtremaFunc returns minimum and maximum from the slice using a function to compare values.
// Panics if the slice is empty or nil.
func ExtremaFunc[T Ordered](sl []T, f LessFunc[T]) (min, max T) {
	if len(sl) == 0 {
		panic("slice is empty or nil")
	}
	min = sl[0]
	max = sl[0]
	for _, v := range sl[1:] {
		if f(v, min) {
			min = v
		} else if f(max, v) {
			max = v
		}
	}
	return
}

// Sort sorts the slice.
func Sort[T Ordered](sl []T) {
	SortFunc(sl, false, lessFn[T])
}

// SortFunc sorts the slice using a function to compare values.
func SortFunc[T any](sl []T, stable bool, f LessFunc[T]) {
	if stable {
		sort.SliceStable(sl, func(i, j int) bool {
			return f(sl[i], sl[j])
		})
	} else {
		sort.Slice(sl, func(i, j int) bool {
			return f(sl[i], sl[j])
		})
	}
}

// SortClone returns a sorted clone of the slice.
func SortClone[T Ordered](sl []T) []T {
	return SortCloneFunc(sl, false, lessFn[T])
}

// SortCloneFunc returns a sorted clone of the slice using a function to compare values.
func SortCloneFunc[T any](sl []T, stable bool, f LessFunc[T]) []T {
	result := Clone(sl)
	SortFunc(result, stable, f)
	return result
}

// IsSorted returns true if the slice is sorted.
func IsSorted[T Ordered](sl []T) bool {
	return IsSortedFunc(sl, lessFn[T])
}

// IsSortedFunc returns true if the slice is sorted using a function to compare values.
func IsSortedFunc[T any](sl []T, f LessFunc[T]) bool {
	return sort.SliceIsSorted(sl, func(i, j int) bool {
		return f(sl[i], sl[j])
	})
}

// Search does a binary search and returns idx and found.
// If found is true, idx is the first index at which v was found.
// If found is false, idx is the index at which v must be inserted to keep sl sorted.
// The slice sl must be sorted in ascending order.
func Search[T Ordered](sl []T, v T) (idx int, found bool) {
	l := len(sl)
	i := sort.Search(l, func(i int) bool { return sl[i] >= v })
	if i < l && sl[i] == v {
		return i, true
	} else {
		return i, false
	}
}

// SearchFunc does a binary search using a function to compare values.
// If found is true, idx is the first index at which v was found.
// If found is false, idx is the index at which v must be inserted to keep sl sorted.
// The slice sl must be sorted in ascending order according to function f.
func SearchFunc[T any](sl []T, v T, f LessFunc[T]) (int, bool) {
	l := len(sl)
	i := sort.Search(l, func(i int) bool { return !f(sl[i], v) })
	if i < l && !(f(sl[i], v) || f(v, sl[i])) {
		return i, true
	} else {
		return i, false
	}
}
