package slices

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	var tests = []struct {
		sl   []int
		v    int
		want bool
	}{
		{[]int{}, 1, false},
		{[]int{1}, 1, true},
		{[]int{1, 2}, 2, true},
		{[]int{1, 2}, 3, false},
	}
	for i, test := range tests {
		if got := Contains(test.sl, test.v); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}

func TestEqual(t *testing.T) {
	var tests = []struct {
		sl1, sl2 []int
		want     bool
	}{
		{nil, nil, true},
		{[]int{}, nil, true},
		{[]int{1}, nil, false},
		{[]int{1}, []int{1}, true},
		{[]int{1}, []int{2}, false},
		{[]int{1, 2}, []int{1, 2}, true},
		{[]int{1, 2}, []int{2, 1}, false},
		{[]int{1, 2}, []int{1, 2, 3}, false},
	}
	for i, test := range tests {
		got1 := Equal(test.sl1, test.sl2)
		got2 := Equal(test.sl2, test.sl1)
		if (got1 != test.want) || (got2 != test.want) {
			t.Errorf("%d: got %t and %t, want %t", i, got1, got2, test.want)
		}
	}
}

func TestFind(t *testing.T) {
	var tests = []struct {
		sl   []int
		v    int
		want int
	}{
		{[]int{}, 1, -1},
		{[]int{1}, 1, 0},
		{[]int{1, 2}, 2, 1},
		{[]int{1, 2}, 3, -1},
	}
	for i, test := range tests {
		if got := Find(test.sl, test.v); got != test.want {
			t.Errorf("%d: got %d, want %d", i, got, test.want)
		}
	}
}

func TestFindAll(t *testing.T) {
	var tests = []struct {
		sl   []int
		v    int
		want []int
	}{
		{nil, 1, nil},
		{[]int{}, 1, []int{}},
		{[]int{1}, 1, []int{0}},
		{[]int{2}, 1, []int{}},
		{[]int{1, 2, 1}, 1, []int{0, 2}},
	}
	for i, test := range tests {
		if got := FindAll(test.sl, test.v); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v,, want %#v", i, got, test.want)
		}
	}
}

func TestUnique(t *testing.T) {
	var tests = []struct {
		sl, want []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1, 2}, []int{2, 1}},
	}
	for i, test := range tests {
		if got := Unique(test.sl); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestUniqueFunc(t *testing.T) {
	var tests = []struct {
		sl, want []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1, 2}, []int{2, 1}},
	}
	for i, test := range tests {
		if got := UniqueFunc(test.sl, cmpFn[int]); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}
