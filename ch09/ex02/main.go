package main

import (
	"fmt"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch09/ex02/popcount"
)

func main() {
	fmt.Println(popcount.PopCount(0x1))
	fmt.Println(popcount.PopCount(0x11))
	fmt.Println(popcount.PopCount(0x110))
}