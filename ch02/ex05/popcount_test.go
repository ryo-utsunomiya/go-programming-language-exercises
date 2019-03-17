package popcount

import (
	"testing"
)

func TestPopCount1(t *testing.T) {
	got := PopCount1(uint64(6))
	want := 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPopCount2(t *testing.T) {
	got := PopCount2(uint64(6))
	want := 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPopCount3(t *testing.T) {
	got := PopCount3(uint64(6))
	want := 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPopCount4(t *testing.T) {
	got := PopCount4(uint64(6))
	want := 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount1(uint64(i))
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(uint64(i))
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(uint64(i))
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount4(uint64(i))
	}
}
