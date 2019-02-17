package main

import (
	"fmt"
	"io"
	"os"
)

func Echo(out io.Writer, args []string) {
	for i, arg := range args {
		_, err := fmt.Fprintf(out, "%d: %s\n", i, arg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func main() {
	Echo(os.Stdout, os.Args[1:])
}
