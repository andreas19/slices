package slices

import (
	"math/rand"
)

// Shuffle pseudo-randomizes the order of the elements in the slice.
func Shuffle[T any](sl []T) {
	rand.Shuffle(len(sl), func(i, j int) {
		sl[i], sl[j] = sl[j], sl[i]
	})
}

// RandomInt returns a slice of given length with pseudo-random values in the interval [0,n).
// n == 0 means no upper limit. It panics when n < 0
func RandomInt(length, n int) []int {
	var randFunc func() int
	if n == 0 {
		randFunc = rand.Int
	} else {
		randFunc = func() int { return rand.Intn(n) }
	}
	return CreateFunc(length, randFunc)
}

// RandomFloat returns a slice of given length with pseudo-random values in the interval [0.0,1.0).
func RandomFloat(length int) []float64 {
	return CreateFunc(length, rand.Float64)
}
