package slices

// A MapFunc maps a value to another value which may have a different type.
type MapFunc[T, U any] func(x T) U

// Map maps a values in a slice using a mapping function.
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
