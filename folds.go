package slices

import (
	"reflect"
)

// Reduce function
type ReduceFunc[T, U any] func(acc U, v T) U

// Reduce reduces sl to a single value by accumulating the values in sl using function f.
// The first argument to f is the accumulated value (starting with init).
func Reduce[T, U any](sl []T, init U, f ReduceFunc[T, U]) U {
	result := init
	for _, v := range sl {
		result = f(result, v)
	}
	return result
}

var boolSliceType = reflect.SliceOf(reflect.TypeOf(false))

// All returns true if for all values in sl the function pred returns true.
// Returns true for an empty slice.
// If pred is nil, type T must be bool.
func All[T any](sl []T, pred Predicate[T]) bool {
	if pred == nil {
		if boolSliceType != reflect.TypeOf(sl) {
			panic("slice must be of type []bool if pred is nil")
		}
		for _, v := range sl {
			if !reflect.ValueOf(v).Bool() {
				return false
			}
		}
	} else {
		for _, v := range sl {
			if !pred(v) {
				return false
			}
		}
	}
	return true
}

// Any returns true if for any value in sl the function pred returns true.
// Returns false for an empty slice.
// If pred is nil, type T must be bool.
func Any[T any](sl []T, pred Predicate[T]) bool {
	if pred == nil {
		if boolSliceType != reflect.TypeOf(sl) {
			panic("slice must be of type []bool if pred is nil")
		}
		for _, v := range sl {
			if reflect.ValueOf(v).Bool() {
				return true
			}
		}
	} else {
		for _, v := range sl {
			if pred(v) {
				return true
			}
		}
	}
	return false
}
