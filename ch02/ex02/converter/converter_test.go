package converter

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

func TestMToF(t *testing.T) {
	got := MToF(Meter(1)).String()
	want := "3.2808ft"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFToM(t *testing.T) {
	got := FToM(Feet(1)).String()
	want := "0.30480370641307m"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestKToP(t *testing.T) {
	got := KToP(KiloGram(1)).String()
	want := "2.20462£"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPToK(t *testing.T) {
	got := PToK(Pound(1)).String()
	want := "0.45359290943563974Kg"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
