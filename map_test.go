package slices

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	var tests = []struct {
		sl, want []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 4}},
	}
	f := func(x int) int { return x * x }
	for i, test := range tests {
		if got := Map(test.sl, f); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}
