package slices

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	var tests = []struct {
		sl, vs []int
		idx    int
		want   []int
	}{
		{nil, nil, 0, nil},
		{nil, []int{}, 0, nil},
		{nil, []int{1}, 0, []int{1}},
		{[]int{}, []int{}, 0, []int{}},
		{[]int{}, []int{1}, 0, []int{1}},
		{[]int{0, 1, 2, 3}, []int{}, 0, []int{0, 1, 2, 3}},
		{[]int{0, 1, 2, 3}[:2], []int{4}, 0, []int{4, 0, 1}},
		{[]int{0, 1, 2, 3}[:2], []int{4}, 1, []int{0, 4, 1}},
		{[]int{0, 1, 2, 3}[:2], []int{4}, 2, []int{0, 1, 4}},
		{[]int{0, 1, 2, 3}[:2], []int{4, 5}, 1, []int{0, 4, 5, 1}},
		{[]int{0, 1, 2, 3}[:2], []int{4, 5, 6}, 0, []int{4, 5, 6, 0, 1}},
		{[]int{0, 1, 2, 3}[:2], []int{4, 5, 6}, 1, []int{0, 4, 5, 6, 1}},
		{[]int{0, 1, 2, 3}[:2], []int{4, 5, 6}, 2, []int{0, 1, 4, 5, 6}},
	}
	for i, test := range tests {
		if got := Insert(test.sl, test.idx, test.vs...); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestDelete(t *testing.T) {
	var tests = []struct {
		sl, idxs, want []int
	}{
		{nil, nil, nil},
		{[]int{}, nil, []int{}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{1}, []int{0}, []int{}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{0}, []int{2, 3}},
		{[]int{1, 2, 3}, []int{1}, []int{1, 3}},
		{[]int{1, 2, 3}, []int{2}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{0, 1}, []int{3}},
		{[]int{1, 2, 3}, []int{1, 2}, []int{1}},
		{[]int{1, 2, 3}, []int{0, 2}, []int{2}},
		{[]int{1, 2, 3}, []int{0, 1, 2}, []int{}},
	}
	for i, test := range tests {
		if got := Delete(test.sl, test.idxs...); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestDeleteRange(t *testing.T) {
	var tests = []struct {
		sl            []int
		start, length int
		want          []int
	}{
		{[]int{1, 2, 3}, 0, 1, []int{2, 3}},
		{[]int{1, 2, 3}, 0, 2, []int{3}},
		{[]int{1, 2, 3}, 0, 3, []int{}},
		{[]int{1, 2, 3}, 0, 4, []int{}},
		{[]int{1, 2, 3}, 1, 1, []int{1, 3}},
		{[]int{1, 2, 3}, 1, 2, []int{1}},
	}
	for i, test := range tests {
		if got := DeleteRange(test.sl, test.start, test.length); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestDeletePred(t *testing.T) {
	sl := []int{1, 2, 3, 4}
	want := []int{1, 3}
	p := func(v int) bool { return v%2 == 0 }
	if got := DeletePred(sl, p); !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestFill(t *testing.T) {
	var tests = []struct {
		sl   []int
		v    int
		want []int
	}{
		{nil, 0, nil},
		{[]int{}, 0, []int{}},
		{[]int{1}, 0, []int{0}},
		{[]int{1, 2}, 0, []int{0, 0}},
		{[]int{1, 2, 3, 4}[:2], 0, []int{0, 0, 0, 0}},
	}
	for i, test := range tests {
		Fill(test.sl, test.v)
		if c := cap(test.sl); !reflect.DeepEqual(test.sl[:c], test.want) ||
			!reflect.DeepEqual(test.sl, test.want[:len(test.sl)]) {
			t.Errorf("%d: got %#v, want %#v", i, test.sl[:c], test.want)
		}
	}
}
