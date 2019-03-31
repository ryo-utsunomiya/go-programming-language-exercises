package main

import (
	"crypto/sha256"
	"fmt"
)

func CountSha256HashBitDiff(a, b [sha256.Size]byte) int {
	var cnt int

	for i := 0; i < 32; i++ {
		if a[i] != b[i] {
			cnt++
		}
	}

	return cnt
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%d\n", CountSha256HashBitDiff(c1, c2))
}
