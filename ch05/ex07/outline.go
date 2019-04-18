package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		if err := outline(os.Stdout, url); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}
}

func outline(w io.Writer, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	o := ForEachNode(doc, startElement, endElement)
	if _, err := w.Write([]byte(o)); err != nil {
		return err
	}

	return nil
}

func ForEachNode(n *html.Node, pre, post func(n *html.Node) string) string {
	result := ""

	if pre != nil {
		result += pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result += ForEachNode(c, pre, post)
	}

	if post != nil {
		result += post(n)
	}

	return result
}

var depth int

func startElement(n *html.Node) string {
	result := ""

	switch n.Type {
	case html.ElementNode:
		var attr string
		for _, a := range n.Attr {
			attr += fmt.Sprintf(" %s=%q", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			result += fmt.Sprintf("%*s<%s%s/>\n", depth*2, "", n.Data, attr)
		} else {
			result += fmt.Sprintf("%*s<%s%s>\n", depth*2, "", n.Data, attr)
			depth++
		}
	case html.TextNode:
		s := strings.TrimSpace(n.Data)
		if len(s) > 0 {
			result += fmt.Sprintf("%*s%s\n", depth*2, "", s)
		}
	case html.CommentNode:
		result += fmt.Sprintf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	}

	return result
}

func endElement(n *html.Node) string {
	result := ""
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		result += fmt.Sprintf("%*s</%s>\n", depth*2, "", n.Data)
	}
	return result
}
