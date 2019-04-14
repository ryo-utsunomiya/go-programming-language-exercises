package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	TextContent(os.Stdout, doc)
}

func TextContent(w io.Writer, n *html.Node) {
	if n.Type == html.TextNode &&
		n.Parent.Data != "script" &&
		n.Parent.Data != "style" &&
		strings.TrimSpace(n.Data) != "" {
		fmt.Fprintln(w, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		TextContent(w, c)
	}
}
