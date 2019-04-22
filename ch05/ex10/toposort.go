package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range TopoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func TopoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visit func(string)
	visit = func(item string) {
		if !seen[item] {
			seen[item] = true
			for _, dep := range m[item] {
				visit(dep)
			}
			order = append(order, item)
		}
	}

	for item := range m {
		visit(item)
	}

	return order
}
