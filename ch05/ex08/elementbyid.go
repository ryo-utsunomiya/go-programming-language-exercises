package main

import (
	"golang.org/x/net/html"
)

func ElementById(doc *html.Node, id string) *html.Node {
	return ForEachNode(doc, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					return false
				}
			}
		}
		return true
	}, nil)
}

func ForEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if !pre(n) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if cn := ForEachNode(c, pre, post); cn != nil {
			return cn
		}
	}

	if post != nil {
		if !post(n) {
			return n
		}
	}

	return nil
}
