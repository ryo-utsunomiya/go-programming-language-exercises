package main

import "testing"

func TestReverse(t *testing.T) {
	want := [3]int{3, 2, 1}
	got := [3]int{1, 2, 3}
	Reverse(&got)

	for i := 0; i < 3; i++ {
		if want[i] != got[i] {
			t.Errorf("want: %d, got: %d", want[i], got[i])
		}
	}
}
