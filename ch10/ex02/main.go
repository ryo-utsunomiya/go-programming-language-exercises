package main

import (
	"fmt"
	"os"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch10/ex02/reader"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: reader foo.zip")
		os.Exit(1)
	}

	err := reader.Read(os.Args[1], os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
