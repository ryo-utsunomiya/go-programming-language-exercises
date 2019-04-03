package main

import "fmt"

func Reverse(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

func main() {
	s := []byte("abc")
	Reverse(s)
	fmt.Println(string(s))
}