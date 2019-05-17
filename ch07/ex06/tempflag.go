package main

import (
	"flag"
	"fmt"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch07/ex06/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
