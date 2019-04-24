package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (n int) {
	defer func() {
		recover()
		n = 42
	}()
	panic(nil)
}
