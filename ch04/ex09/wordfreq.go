package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	for _, arg := range os.Args[1:] {
		wordfreq(arg)
	}
}

func wordfreq(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	defer file.Close()

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	words := make(map[string]int)

	for input.Scan() {
		words[input.Text()]++
	}

	type KV struct {
		k string
		v int
	}
	var sorted []KV
	for k, v := range words {
		sorted = append(sorted, KV{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].v > sorted[j].v
	})

	for _, kv := range sorted {
		fmt.Printf("%q\t%d\n", kv.k, kv.v)
	}
}
