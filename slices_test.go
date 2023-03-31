package slices

import (
	"reflect"
	"testing"
)

func TestClone(t *testing.T) {
	var tests = [][]int{nil, {}, {1}, {1, 2}}
	for i, test := range tests {
		if got := Clone(test); reflect.DeepEqual(got, test) {
			if len(got) > 0 && &got[0] == &test[0] {
				t.Errorf("%d: uses same array", i)
			}
		} else {
			t.Errorf("%d: got %#v, want %#v", i, got, test)
		}
	}
}

func TestConcat(t *testing.T) {
	if got := Concat[[]int](); got != nil {
		t.Errorf("got %#v, want nil slice", got)
	}
	want := []int{1, 2, 3}
	if got := Concat([]int{1}, []int{}, nil, []int{2, 3}); !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestFilter(t *testing.T) {
	f := func(v int) bool { return v%2 == 0 }
	var tests = []struct{ sl, want []int }{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{}},
		{[]int{2}, []int{2}},
		{[]int{1, 2}, []int{2}},
		{[]int{1, 3}, []int{}},
		{[]int{1, 2, 3, 4}, []int{2, 4}},
	}
	for i, test := range tests {
		if got := Filter(test.sl, f); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestRepeat(t *testing.T) {
	var tests = []struct {
		v, n int
		want []int
	}{
		{1, 0, []int{}},
		{1, 1, []int{1}},
		{1, 2, []int{1, 1}},
	}
	for i, test := range tests {
		if got := Repeat(test.v, test.n); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestRepeatNegN(t *testing.T) {
	defer func() { _ = recover() }()
	_ = Repeat(1, -1)
	t.Error("did not panic")
}

func TestReverse(t *testing.T) {
	var tests = []struct{ sl, want []int }{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}
	for i, test := range tests {
		if Reverse(test.sl); !reflect.DeepEqual(test.sl, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, test.sl, test.want)
		}
	}
}

func TestReverseClone(t *testing.T) {
	var tests = []struct{ sl, want []int }{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}
	for i, test := range tests {
		if got := ReverseClone(test.sl); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}
