package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for name, cnt := range MapElementNum(nil, doc) {
		fmt.Printf("%s:%d\n", name, cnt)
	}
}

func MapElementNum(m map[string]int, n *html.Node) map[string]int {
	if m == nil {
		m = make(map[string]int)
	}

	if n.Type == html.ElementNode {
		m[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		MapElementNum(m, c)
	}

	return m
}
