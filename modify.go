package slices

// Insert inserts values at index idx and returns the modified slice.
// If there is enough unused capacity, the original slice will be modified.
func Insert[T any](sl []T, idx int, v ...T) []T {
	if sl == nil && len(v) == 0 {
		return nil
	}
	var result []T
	if len(v) <= cap(sl)-len(sl) {
		result = sl[:len(sl)+len(v)]
	} else {
		result = make([]T, len(sl)+len(v))
		copy(result, sl[:idx])
	}
	copy(result[idx+len(v):], sl[idx:])
	copy(result[idx:], v)
	return result
}
