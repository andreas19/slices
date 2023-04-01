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
