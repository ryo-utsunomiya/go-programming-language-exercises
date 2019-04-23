package main

import (
	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, tagName string) []*html.Node {
	elements := make([]*html.Node, 0)

	ForEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tagName {
			elements = append(elements, n)
		}
	}, nil)

	return elements
}

func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
