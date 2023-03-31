package slices

import (
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	var tests = []struct {
		sl   []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 1},
	}
	for i, test := range tests {
		if got := Min(test.sl); got != test.want {
			t.Errorf("%d: got %d, want %d", i, got, test.want)
		}
	}
}

func TestMinNil(t *testing.T) {
	defer func() { _ = recover() }()
	_ = Min[int](nil)
	t.Error("did not panic")
}

func TestMinEmpty(t *testing.T) {
	defer func() { _ = recover() }()
	_ = Min([]int{})
	t.Error("did not panic")
}

func TestMax(t *testing.T) {
	var tests = []struct {
		sl   []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 1}, 2},
	}
	for i, test := range tests {
		if got := Max(test.sl); got != test.want {
			t.Errorf("%d: got %d, want %d", i, got, test.want)
		}
	}
}

func TestMaxNil(t *testing.T) {
	defer func() { _ = recover() }()
	_ = Max[int](nil)
	t.Error("did not panic")
}

func TestMaxEmpty(t *testing.T) {
	defer func() { _ = recover() }()
	_ = Max([]int{})
	t.Error("did not panic")
}

func TestSort(t *testing.T) {
	var tests = []struct {
		sl, want []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
	}
	for i, test := range tests {
		if Sort(test.sl); !reflect.DeepEqual(test.sl, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, test.sl, test.want)
		}
	}
}

func TestSortClone(t *testing.T) {
	var tests = []struct {
		sl, want []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
	}
	for i, test := range tests {
		if got := SortClone(test.sl); reflect.DeepEqual(got, test.want) {
			if len(got) > 0 && &got[0] == &test.sl[0] {
				t.Errorf("%d: uses same array", i)
			}
		} else {
			t.Errorf("%d: got %#v, want %#v", i, test.sl, test.want)
		}
	}
}

func TestIsSorted(t *testing.T) {
	var tests = []struct {
		sl   []int
		want bool
	}{
		{nil, true},
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 2}, true},
		{[]int{2, 1}, false},
	}
	for i, test := range tests {
		if got := IsSorted(test.sl); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}
