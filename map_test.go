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

func TestMap2(t *testing.T) {
	var tests = []struct {
		sl1, sl2, want []int
	}{
		{nil, nil, nil},
		{nil, []int{}, nil},
		{[]int{}, nil, nil},
		{[]int{}, []int{1}, []int{}},
		{[]int{1}, []int{1}, []int{2}},
		{[]int{1, 2}, []int{3, 4}, []int{4, 6}},
	}
	f := func(x, y int) int { return x + y }
	for i, test := range tests {
		if got := Map2(test.sl1, test.sl2, f); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestMap3(t *testing.T) {
	var tests = []struct {
		sl1, sl2, sl3, want []int
	}{
		{nil, nil, nil, nil},
		{nil, []int{}, []int{}, nil},
		{[]int{}, nil, []int{}, nil},
		{[]int{}, []int{1}, []int{}, []int{}},
		{[]int{1}, []int{1}, []int{1}, []int{3}},
		{[]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{9, 12}},
	}
	f := func(x, y, z int) int { return x + y + z }
	for i, test := range tests {
		if got := Map3(test.sl1, test.sl2, test.sl3, f); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}
