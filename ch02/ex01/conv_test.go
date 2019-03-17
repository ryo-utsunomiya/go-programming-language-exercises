package tempconv

import (
	"testing"
)

func TestFToC(t *testing.T) {
	got := FToC(Fahrenheit(451)).String()
	want := "232.77777777777777℃"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCToF(t *testing.T) {
	got := CToF(Celsius(100)).String()
	want := "212℉"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCToK(t *testing.T) {
	got := CToK(AbsoluteZeroC).String()
	want := "0K"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestKToC(t *testing.T) {
	got := KToC(AbsoluteZeroK).String()
	want := "-273.15℃"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
