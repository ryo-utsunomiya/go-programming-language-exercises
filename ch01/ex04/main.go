package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	Dup(os.Stdout, os.Args[1:])
}

func Dup(out io.Writer, files []string) {
	counts := make(map[string]int)

	for _, filename := range files {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("dup: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[filename+"\t"+line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			_, err := fmt.Fprintf(out, "%d\t%s\n", n, line)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}
