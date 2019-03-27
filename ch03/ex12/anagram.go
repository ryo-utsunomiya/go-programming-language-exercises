package main

import (
	"fmt"
	"sort"
	"strings"
)

func IsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	as := strings.Split(a, "")
	bs := strings.Split(b, "")
	sort.Slice(as, func(i, j int) bool {
		return as[i] < as[j]
	})
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})

	for i, _ := range as {
		if as[i] != bs[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(IsAnagram("however", "whoever"))
	fmt.Println(IsAnagram("nowhere", "now here"))
}
