package slices

// Predicate function type
type Predicate[T any] func(v T) bool

func Clone[T any](sl []T) []T {
	if sl == nil {
		return nil
	}
	result := make([]T, len(sl))
	copy(result, sl)
	return result
}

// Compact sets the capacity equal to the length of the slice.
func Compact[T any](sl []T) []T {
	return sl[:len(sl):len(sl)]
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

// Flatten returns a new slice with one level of nesting removed from a slice of slices.
func Flatten[T any](sl [][]T) []T {
	return Concat(sl...)
}

// Repeat returns a slice with n repetitions of the given value.
// Pancis when n < 0
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
