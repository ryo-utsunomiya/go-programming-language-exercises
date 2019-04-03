package main

import "fmt"

func RemoveDuplicatedAdjacentString(xs []string) []string {
	result := make([]string, 0)
	for i, s := range xs {
		if i+1 < len(xs) && s == xs[i+1] {
			continue
		}
		result = append(result, s)
	}
	return result
}

func main() {
	s := []string{"foo", "foo", "bar"}
	fmt.Println(RemoveDuplicatedAdjacentString(s))
}
