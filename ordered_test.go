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
	for _, test := range tests {
		if got := Min(test.sl); got != test.want {
			t.Errorf("got %d, want %d", got, test.want)
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
	for _, test := range tests {
		if got := Max(test.sl); got != test.want {
			t.Errorf("got %d, want %d", got, test.want)
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
	for _, test := range tests {
		if Sort(test.sl); !reflect.DeepEqual(test.sl, test.want) {
			t.Errorf("got %#v, want %#v", test.sl, test.want)
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
	for _, test := range tests {
		if got := SortClone(test.sl); reflect.DeepEqual(got, test.want) {
			if len(got) > 0 && &got[0] == &test.sl[0] {
				t.Error("uses same array")
			}
		} else {
			t.Errorf("got %#v, want %#v", test.sl, test.want)
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
	for _, test := range tests {
		if got := IsSorted(test.sl); got != test.want {
			t.Errorf("got %t, want %t", got, test.want)
		}
	}
}
