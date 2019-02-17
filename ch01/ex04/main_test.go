package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	out := new(bytes.Buffer)
	Dup(out, []string{"sample1.txt"})

	got := out.String()
	want := `2	sample1.txt	foo
3	sample1.txt	bar
`
	if got != want {
		t.Errorf("got\n %v\n want\n %v", got, want)
	}
}
