package main

import "testing"

var commaTests = []struct {
	a, b string
	want bool
}{
	{"", "", true},
	{"aba", "baa", true},
	{"aaa", "baa", false},
	{"however", "whoever", true},
	{"nowhere", "now here", false},
	{"こんにちは", "はちにんこ", true},
	{"こんにちは", "こんばんは", false},
}

func TestIsAnagram(t *testing.T) {
	for _, tt := range commaTests {
		t.Run(tt.a+tt.b, func(t *testing.T) {
			got := IsAnagram(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %v, want: %v\n", got, tt.want)
			}
		})
	}
}
