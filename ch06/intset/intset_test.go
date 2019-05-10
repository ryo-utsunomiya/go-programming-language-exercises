package intset

import (
	"testing"
)

func TestIntSet_Add(t *testing.T) {
	var x IntSet
	x.Add(1)

	want := []uint{2}
	got := x.words

	if got[0] != want[0] {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestIntSet_Has(t *testing.T) {
	var x IntSet
	x.Add(1)

	if !x.Has(1) {
		t.Errorf("want: %v, got: %v", true, x.Has(1))
	}
	if x.Has(2) {
		t.Errorf("want: %v, got: %v", false, x.Has(1))
	}
}
