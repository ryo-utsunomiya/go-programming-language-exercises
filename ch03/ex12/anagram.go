package main

import (
	"fmt"
	"strings"
)

func IsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	tmp := b
	for _, r := range a {
		if !strings.ContainsRune(tmp, r) {
			return false
		}
		tmp = strings.Replace(tmp, string(r), "", 1)
	}

	return tmp == ""
}

func main() {
	fmt.Println(IsAnagram("however", "whoever"))
	fmt.Println(IsAnagram("nowhere", "now here"))
}
