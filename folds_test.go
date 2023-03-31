package slices

import "testing"

func TestReduce(t *testing.T) {
	var tests = []struct {
		sl         []int
		init, want int
	}{
		{nil, 1, 1},
		{[]int{}, 1, 1},
		{[]int{1}, 1, 2},
		{[]int{1, 2}, 1, 4},
	}
	f := func(x, y int) int { return x + y }
	for i, test := range tests {
		if got := Reduce(test.sl, test.init, f); got != test.want {
			t.Errorf("%d: got %d, want %d", i, got, test.want)
		}
	}
}

func TestAll(t *testing.T) {
	var tests = []struct {
		sl   []int
		want bool
	}{
		{nil, true},
		{[]int{}, true},
		{[]int{1, 2, 4}, false},
		{[]int{2, 4}, true},
	}
	p := func(v int) bool { return v%2 == 0 }
	for i, test := range tests {
		if got := All(test.sl, p); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}

func TestAllBool(t *testing.T) {
	var tests = []struct {
		sl   []bool
		want bool
	}{
		{nil, true},
		{[]bool{}, true},
		{[]bool{true, true}, true},
		{[]bool{true, false}, false},
	}
	for i, test := range tests {
		if got := All(test.sl, nil); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}

func TestAllNilPred(t *testing.T) {
	defer func() { _ = recover() }()
	All([]int{}, nil)
	t.Error("did not panic")
}

func TestAny(t *testing.T) {
	var tests = []struct {
		sl   []int
		want bool
	}{
		{nil, false},
		{[]int{}, false},
		{[]int{1, 3}, false},
		{[]int{1, 2, 4}, true},
		{[]int{2, 4}, true},
	}
	p := func(v int) bool { return v%2 == 0 }
	for i, test := range tests {
		if got := Any(test.sl, p); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}

func TestAnyBool(t *testing.T) {
	var tests = []struct {
		sl   []bool
		want bool
	}{
		{nil, false},
		{[]bool{}, false},
		{[]bool{false, false}, false},
		{[]bool{true, true}, true},
		{[]bool{true, false}, true},
	}
	for i, test := range tests {
		if got := Any(test.sl, nil); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}

func TestAnyNilPred(t *testing.T) {
	defer func() { _ = recover() }()
	Any([]int{}, nil)
	t.Error("did not panic")
}
