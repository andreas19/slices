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

func TestExtrema(t *testing.T) {
	var tests = []struct {
		sl                 []int
		want_min, want_max int
	}{
		{[]int{1}, 1, 1},
		{[]int{1, 1}, 1, 1},
		{[]int{1, 2}, 1, 2},
		{[]int{1, 2, 3}, 1, 3},
		{[]int{3, 2, 1}, 1, 3},
	}
	for i, test := range tests {
		if got_min, got_max := Extrema(test.sl); got_min != test.want_min || got_max != test.want_max {
			t.Errorf("%d: got %d and %d, want %d and %d", i, got_min, got_max, test.want_min, test.want_max)
		}
	}
}

func TestExtremaNil(t *testing.T) {
	defer func() { _ = recover() }()
	_, _ = Extrema[int](nil)
	t.Error("did not panic")
}

func TestExtremaEmpty(t *testing.T) {
	defer func() { _ = recover() }()
	_, _ = Extrema([]int{})
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

func TestSearch(t *testing.T) {
	var tests = []struct {
		sl         []int
		v          int
		want_idx   int
		want_found bool
	}{
		{nil, 1, 0, false},
		{[]int{}, 1, 0, false},
		{[]int{1}, 0, 0, false},
		{[]int{1}, 1, 0, true},
		{[]int{1}, 2, 1, false},
		{[]int{1, 3}, 2, 1, false},
	}
	for i, test := range tests {
		if got_idx, got_found := Search(test.sl, test.v); got_idx != test.want_idx || got_found != test.want_found {
			t.Errorf("%d: got %d, %t, want %d, %t", i, got_idx, got_found, test.want_idx, test.want_found)
		}
	}
}

func TestSearchFunc(t *testing.T) {
	var tests = []struct {
		sl         []int
		v          int
		want_idx   int
		want_found bool
	}{
		{nil, 1, 0, false},
		{[]int{}, 1, 0, false},
		{[]int{1}, 0, 0, false},
		{[]int{1}, 1, 0, true},
		{[]int{1}, 2, 1, false},
		{[]int{1, 3}, 2, 1, false},
	}
	for i, test := range tests {
		if got_idx, got_found := SearchFunc(test.sl, test.v, lessFn[int]); got_idx != test.want_idx || got_found != test.want_found {
			t.Errorf("%d: got %d, %t, want %d, %t", i, got_idx, got_found, test.want_idx, test.want_found)
		}
	}
}
