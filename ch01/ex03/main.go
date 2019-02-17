package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Echo1(out io.Writer, args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	_, err := fmt.Fprintln(out, s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Echo3(out io.Writer, args []string) {
	_, err := fmt.Fprintln(out, strings.Join(args, " "))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
