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

// Delete deletes the values at the given indexes and returns the modified slice.
func Delete[T any](sl []T, idx ...int) []T {
	if sl == nil {
		return nil
	}
	if len(idx) == 0 {
		return sl
	}
	idxs := Unique(idx)
	SortFunc(idxs, false, func(v1, v2 int) bool { return v2 < v1 })
	li := len(sl) - 1
	for _, i := range idxs {
		copy(sl[i:], sl[i+1:])
		var v T
		sl[li] = v
		li--
	}
	return sl[:len(sl)-len(idxs)]
}

// DeleteRange deletes length values beginning at start and returns the modified slice.
func DeleteRange[T any](sl []T, start, length int) []T {
	end := start + length
	if l := len(sl); end > l {
		end = l
	}
	return Delete(sl, Range(start, end, 1)...)
}

// DeletePred deletes the values for which pred returns true and returns the modified slice.
func DeletePred[T any](sl []T, pred Predicate[T]) []T {
	var idxs []int
	for i, v := range sl {
		if pred(v) {
			idxs = append(idxs, i)
		}
	}
	return Delete(sl, idxs...)
}
