package slices

// Predicate function type
type Predicate[T any] func(v T) bool

// Clone clones the slice.
func Clone[T any](sl []T) []T {
	if sl == nil {
		return nil
	}
	result := make([]T, len(sl))
	copy(result, sl)
	return result
}

// Concat returns the concatenation of all given slices as a new slice.
func Concat[T any](sls ...[]T) []T {
	if sls == nil {
		return nil
	}
	l := 0
	for _, sl := range sls {
		l += len(sl)
	}
	result := make([]T, 0, l)
	for _, sl := range sls {
		result = append(result, sl...)
	}
	return result
}

// Create returns a new slice with the given values.
func Create[T any](vs ...T) []T {
	return vs
}

// Filter returns a new slice with all values for which pred returns true.
func Filter[T any](sl []T, pred Predicate[T]) []T {
	if sl == nil {
		return nil
	}
	result := []T{}
	for _, v := range sl {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterIndex returns a new slice with all values for which pred(index) returns true.
func FilterIndex[T any](sl []T, pred Predicate[int]) []T {
	if sl == nil {
		return nil
	}
	result := []T{}
	for i, v := range sl {
		if pred(i) {
			result = append(result, v)
		}
	}
	return result
}

// Flatten returns a new slice with one level of nesting removed from a slice of slices.
func Flatten[T any](sl [][]T) []T {
	return Concat(sl...)
}

// Repeat returns a slice with n repetitions of the given value.
// Pancis if n < 0.
func Repeat[T any](v T, n int) []T {
	if n < 0 {
		panic("n must be >= 0")
	}
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = v
	}
	return result
}

// Reverse reverses the order of values in the slice.
func Reverse[T any](sl []T) {
	ri := len(sl) - 1
	for li := range sl[:len(sl)/2] {
		sl[li], sl[ri] = sl[ri], sl[li]
		ri--
	}
}

// ReverseClone returns a clone of the slice with the order of values reversed.
func ReverseClone[T any](sl []T) []T {
	result := Clone(sl)
	Reverse(result)
	return result
}

// DropWhile drops values from sl as long as pred returns true and returns
// the rest in a new slice.
func DropWhile[T any](sl []T, pred Predicate[T]) []T {
	cnt := 0
	for _, v := range sl {
		if pred(v) {
			cnt++
		} else {
			break
		}
	}
	return Clone(sl[cnt:])
}

// TakeWhile takes values from sl as long as pred returns true and returns
// them in a new slice.
func TakeWhile[T any](sl []T, pred Predicate[T]) []T {
	cnt := 0
	for _, v := range sl {
		if pred(v) {
			cnt++
		} else {
			break
		}
	}
	return Clone(sl[:cnt])
}

// Get gets the element at index idx. If idx < 0 it computes the index from the end.
// Panics if idx >= len(sl) or the computed index is < 0.
func Get[T any](sl []T, idx int) T {
	if idx < 0 {
		idx = len(sl) + idx
	}
	return sl[idx]
}

// Slice returns a slice from sl. If start or end < 0 it computes the indexes from the end.
// Panics if the (computed) slice bounds are out of range.
func Slice[T any](sl []T, start, end int) []T {
	if start < 0 {
		start = len(sl) + start
	}
	if end < 0 {
		end = len(sl) + end
	}
	return sl[start:end]
}

// Repeat returns a slice with n repetitions of the values in the given slice.
// Pancis if n < 0.
func RepeatSlice[T any](sl []T, n int) []T {
	if n < 0 {
		panic("n must be >= 0")
	}
	if sl == nil {
		return nil
	}
	result := make([]T, 0, n*len(sl))
	for i := 0; i < n; i++ {
		result = append(result, sl...)
	}
	return result
}

// CreateFunc returns a slice of given length with values created by successive calls to fn.
func CreateFunc[T any](length int, fn func() T) []T {
	result := make([]T, length)
	for i := 0; i < length; i++ {
		result[i] = fn()
	}
	return result
}

// Adjust adjusts length and capacity of the given slice. If capacity < length it is set
// to length. Panics if length < 0.
func Adjust[T any](sl []T, length, capacity int) []T {
	if length < 0 {
		panic("length must be >= 0")
	}
	if sl == nil {
		return nil
	}
	if capacity < length {
		capacity = length
	}
	if capacity <= cap(sl) {
		return sl[:length:capacity]
	}
	result := make([]T, capacity)
	copy(result, sl)
	return result[:length]
}
