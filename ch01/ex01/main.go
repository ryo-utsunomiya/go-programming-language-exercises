package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Echo(out io.Writer, args []string) {
	_, err := fmt.Fprintln(out, strings.Join(args, " "))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Echo(os.Stdout, os.Args[0:]);
}
