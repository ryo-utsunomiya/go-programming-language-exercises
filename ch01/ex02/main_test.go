package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	out := new(bytes.Buffer)
	Echo(out, []string{"foo", "bar", "baz"})

	got := out.String()
	want := "0: foo\n1: bar\n2: baz\n"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
