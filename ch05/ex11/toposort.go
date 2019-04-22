package main

import (
	"fmt"
	"log"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

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

const (
	none = iota
	temp
	perm
)

func TopoSort(m map[string][]string) []string {
	var order []string
	mark := make(map[string]int)

	var visit func(string)
	visit = func(item string) {
		if mark[item] == temp {
			log.Print("TopoSort: m has circular references")
		} else if mark[item] == none {
			mark[item] = temp
			for _, dep := range m[item] {
				visit(dep)
			}
			mark[item] = perm
			order = append(order, item)
		}
	}

	for item := range m {
		if mark[item] == none {
			visit(item)
		}
	}

	return order
}
