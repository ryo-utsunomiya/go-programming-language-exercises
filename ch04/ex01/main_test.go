package main

import (
	"crypto/sha256"
	"testing"
)

func TestCountSha256HashBitDiff(t *testing.T) {
	tests := []struct {
		a    [sha256.Size]byte
		b    [sha256.Size]byte
		want int
	}{
		{sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X")), 31},
		{sha256.Sum256([]byte("x")), sha256.Sum256([]byte("x")), 0},
	}

	for _, c := range tests {
		got := CountSha256HashBitDiff(c.a, c.b)
		if c.want != got {
			t.Errorf("want: %d, got: %d", c.want, got)
		}
	}
}
