package main

import (
	"fmt"
)

func main() {
	fmt.Println(implode("/"))
	fmt.Println(implode("/", "a"))
	fmt.Println(implode("/", "a", "b", "c"))
}

func implode(glue string, pieces ...string) string {
	switch len(pieces) {
	case 0:
		return ""
	case 1:
		return pieces[0]
	}

	result := pieces[0]
	for _, s := range pieces[1:] {
		result += glue
		result += s
	}
	return result
}
