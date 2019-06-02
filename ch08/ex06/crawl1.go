package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch05/ex13/links"
)

var tokens = make(chan struct{}, 20)
var depth = 0
var maxDepth = flag.Int("depth", 1, "crawling depth")

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	flag.Parse()

	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- flag.Args() }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++

				go func(link string) {
					if depth >= *maxDepth {
						worklist <- nil
						n = 0
					} else {
						worklist <- crawl(link)
						depth++
					}
				}(link)
			}
		}
	}
}
