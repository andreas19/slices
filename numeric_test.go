package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	var tests = []struct {
		sl   []int
		want int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 3},
	}
	for i, test := range tests {
		if got := Sum(test.sl); got != test.want {
			t.Errorf("%d: got %d, want %d", i, got, test.want)
		}
	}
}

func TestProduct(t *testing.T) {
	var tests = []struct {
		sl   []int
		want int
	}{
		{nil, 1},
		{[]int{}, 1},
		{[]int{2}, 2},
		{[]int{2, 3}, 6},
	}
	for i, test := range tests {
		if got := Product(test.sl); got != test.want {
			t.Errorf("%d: got %d, want %d", i, got, test.want)
		}
	}
}

func TestRange(t *testing.T) {
	var tests = []struct {
		start, end, step int
		want             []int
	}{
		{0, 0, 1, nil},
		{0, 1, 1, []int{0}},
		{0, 4, 1, []int{0, 1, 2, 3}},
		{0, 4, 2, []int{0, 2}},
		{0, 0, -1, nil},
		{0, -1, -1, []int{0}},
		{0, -4, -1, []int{0, -1, -2, -3}},
		{0, -4, -2, []int{0, -2}},
		{0, 1, -1, nil},
		{1, 0, 1, nil},
	}
	for i, test := range tests {
		if got := Range(test.start, test.end, test.step); !reflect.DeepEqual(got, test.want) || cap(got) != len(test.want) {
			t.Errorf("%d: got %#v / %d, want %#v / %d", i, got, cap(got), test.want, len(test.want))
		}
	}
}

func TestRangeStep0(t *testing.T) {
	defer func() { _ = recover() }()
	_ = Range(0, 1, 0)
	t.Error("did not panic")
}
