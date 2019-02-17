package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	out := new(bytes.Buffer)
	Echo(out, []string{"./main", "foo", "bar", "baz"})

	got := out.String()
	want := "./main foo bar baz\n"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
