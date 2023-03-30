package slices

type Float interface {
	~float32 | ~float64
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Number interface {
	Float | Signed | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Sum returns the sum of the values in a slice.
func Sum[T Number](sl []T) T {
	var result T
	for _, v := range sl {
		result += v
	}
	return result
}

// Product returns the product of the values in a slice.
func Product[T Number](sl []T) T {
	var result T = 1
	for _, v := range sl {
		result *= v
	}
	return result
}

// Range returns a slice with numbers in [start, end).
// Will panic if step is 0.
func Range[T Signed | Float](start, end, step T) []T {
	if step == 0 {
		panic("step cannot be 0")
	}
	l := int((end - start) / step)
	if l <= 0 {
		return nil
	}
	result := make([]T, 0, l)
	if step > 0 {
		for v := start; v < end; v += step {
			result = append(result, v)
		}
	} else {
		for v := start; v > end; v += step {
			result = append(result, v)
		}
	}
	return result
}
