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

func TestCompress(t *testing.T) {
	var tests = []struct {
		sl   []int
		sel  []bool
		want []int
	}{
		{nil, nil, nil},
		{[]int{}, []bool{}, []int{}},
		{[]int{1, 2, 3}, []bool{false, true, true}, []int{2, 3}},
		{[]int{1, 2, 3}, []bool{false, true}, []int{2}},
		{[]int{1, 2}, []bool{false, true, true}, []int{2}},
	}
	for i, test := range tests {
		if got := Compress(test.sl, test.sel); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestDropWhile(t *testing.T) {
	var tests = []struct {
		sl   []int
		want []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1}},
		{[]int{2, 2}, []int{}},
		{[]int{2, 3, 2}, []int{3, 2}},
	}
	p := func(x int) bool { return x%2 == 0 }
	for i, test := range tests {
		if got := DropWhile(test.sl, p); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestTakeWhile(t *testing.T) {
	var tests = []struct {
		sl   []int
		want []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1, 2}, []int{}},
		{[]int{2, 1}, []int{2}},
		{[]int{2, 2}, []int{2, 2}},
		{[]int{2, 3, 2}, []int{2}},
	}
	p := func(x int) bool { return x%2 == 0 }
	for i, test := range tests {
		if got := TakeWhile(test.sl, p); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestGet(t *testing.T) {
	sl := []int{1, 2, 3}
	var tests = []struct {
		idx  int
		want int
	}{
		{0, 1},
		{1, 2},
		{2, 3},
		{-1, 3},
		{-2, 2},
		{-3, 1},
	}
	for i, test := range tests {
		if got := Get(sl, test.idx); got != test.want {
			t.Errorf("%d: got %d, want %d", i, got, test.want)
		}
	}
}

func TestGetPanic(t *testing.T) {
	defer func() { _ = recover() }()
	sl := []int{1, 2, 3}
	Get(sl, -4)
	t.Errorf("did not panic")
}

func TestSlice(t *testing.T) {
	sl := []int{1, 2, 3}
	var tests = []struct {
		start, end int
		want       []int
	}{
		{0, 0, []int{}},
		{0, 1, []int{1}},
		{0, 2, []int{1, 2}},
		{0, 3, []int{1, 2, 3}},
		{1, 3, []int{2, 3}},
		{-3, -1, []int{1, 2}},
		{-3, -2, []int{1}},
		{-3, -3, []int{}},
		{-2, -1, []int{2}},
		{-2, -2, []int{}},
		{-1, -1, []int{}},
	}
	for i, test := range tests {
		if got := Slice(sl, test.start, test.end); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestRepeatSlice(t *testing.T) {
	var tests = []struct {
		sl   []int
		n    int
		want []int
	}{
		{nil, 1, nil},
		{[]int{}, 0, []int{}},
		{[]int{}, 1, []int{}},
		{[]int{1}, 0, []int{}},
		{[]int{1}, 1, []int{1}},
		{[]int{1}, 2, []int{1, 1}},
		{[]int{1, 2}, 0, []int{}},
		{[]int{1, 2}, 1, []int{1, 2}},
		{[]int{1, 2}, 2, []int{1, 2, 1, 2}},
	}
	for i, test := range tests {
		if got := RepeatSlice(test.sl, test.n); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestRepeatSliceNegN(t *testing.T) {
	defer func() { _ = recover() }()
	RepeatSlice([]int{1, 2}, -1)
	t.Errorf("did not panic")
}

func TestCreateFunc(t *testing.T) {
	sl := []int{1, 2, 3}
	idx := -1
	want := []int{2, 4}
	got := CreateFunc(2, func() int {
		idx++
		return sl[idx] * 2
	})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestAdjust(t *testing.T) {
	var tests = []struct {
		sl               []int
		length, capacity int
		want_sl          []int
		want_cap         int
	}{
		{nil, 0, 0, nil, 0},
		{[]int{}, 0, 0, []int{}, 0},
		{[]int{}, 0, 1, []int{}, 1},
		{[]int{1, 2, 3, 4, 5}, 0, 0, []int{}, 0},
		{[]int{1, 2, 3, 4, 5}, 0, 1, []int{}, 1},
		{[]int{1, 2, 3, 4, 5}, 0, 2, []int{}, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 0, []int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 2, 2, []int{1, 2}, 2},
		{[]int{1, 2, 3, 4, 5}, 3, 6, []int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4, 5}, 6, 6, []int{1, 2, 3, 4, 5, 0}, 6},
		{[]int{1, 2, 3, 4, 5}, 6, 10, []int{1, 2, 3, 4, 5, 0}, 10},
	}
	for i, test := range tests {
		if got := Adjust(test.sl, test.length, test.capacity); test.want_cap != cap(got) ||
			!reflect.DeepEqual(got, test.want_sl) {
			t.Errorf("%d: got %#v and %d, want %#v and %d", i, got, cap(got),
				test.want_sl, test.want_cap)
		}
	}
}

func TestAdjustNegLength(t *testing.T) {
	defer func() { _ = recover() }()
	Adjust([]int{}, -1, 1)
	t.Error("did not panic")
}
