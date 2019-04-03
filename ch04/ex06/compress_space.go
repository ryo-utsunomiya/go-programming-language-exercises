package main

import (
	"fmt"
	"unicode"
)

func CompressSpace(bytes []byte) []byte {
	result := make([]byte, 0)
	for i, b := range bytes {
		if i+1 < len(bytes) && b == bytes[i+1] && unicode.IsSpace(rune(b)) {
			continue
		}
		result = append(result, b)

	}
	return result
}

func main() {
	fmt.Println(string(CompressSpace([]byte("foo  bar"))))
}