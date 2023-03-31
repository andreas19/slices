package slices

// A MapFunc maps a value to another value which may have a different type.
type MapFunc[T, U any] func(x T) U

// Map maps values in a slice using a mapping function.
func Map[T, U any](sl []T, f MapFunc[T, U]) []U {
	if sl == nil {
		return nil
	}
	result := make([]U, len(sl))
	for i, v := range sl {
		result[i] = f(v)
	}
	return result
}

// A Map2Func maps two values to another value.
type Map2Func[T, U, V any] func(x T, y U) V

// Map2 maps values in two slices using a mapping function.
// It stops when the shortest slice is exhaustet.
func Map2[T, U, V any](sl1 []T, sl2 []U, f Map2Func[T, U, V]) []V {
	if sl1 == nil || sl2 == nil {
		return nil
	}
	length := Min([]int{len(sl1), len(sl2)})
	result := make([]V, length)
	for i := 0; i < length; i++ {
		result[i] = f(sl1[i], sl2[i])
	}
	return result
}

// A Map2Func maps three values to another value.
type Map3Func[T, U, V, W any] func(x T, y U, z V) W

// Map3 maps values in three slices using a mapping function.
// It stops when the shortest slice is exhaustet.
func Map3[T, U, V, W any](sl1 []T, sl2 []U, sl3 []V, f Map3Func[T, U, V, W]) []W {
	if sl1 == nil || sl2 == nil || sl3 == nil {
		return nil
	}
	length := Min([]int{len(sl1), len(sl2), len(sl3)})
	result := make([]W, length)
	for i := 0; i < length; i++ {
		result[i] = f(sl1[i], sl2[i], sl3[i])
	}
	return result
}
