package main

import "testing"

var commaTests = []struct {
	in  string
	out string
}{
	{"123", "123"},
	{"1234", "1,234"},
	{"1234567", "1,234,567"},
}

func TestComma(t *testing.T) {
	for _, tt := range commaTests {
		t.Run(tt.in, func(t *testing.T) {
			if Comma(tt.in) != tt.out {
				t.Errorf("got %s, want: %s\n", tt.in, tt.out)
			}
		})
	}
}
