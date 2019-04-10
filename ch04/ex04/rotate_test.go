package main

import (
	"fmt"
	"testing"
)

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRotate(t *testing.T) {
	tests := []struct {
		s    []int
		n    int
		want []int
	}{
		{
			[]int{0, 1, 2, 3, 4},
			0,
			[]int{0, 1, 2, 3, 4},
		},
		{
			[]int{0, 1, 2, 3, 4},
			1,
			[]int{1, 2, 3, 4, 0},
		},
		{
			[]int{0, 1, 2, 3, 4},
			7,
			[]int{2, 3, 4, 0, 1},
		},
	}

	for _, test := range tests {
		got := Rotate(test.s, test.n)
		if !equals(got, test.want) {
			t.Errorf("case %s: got %v, want %v", fmt.Sprintf("%v", got), test.want, got)
		}
	}
}