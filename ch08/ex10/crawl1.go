package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch08/ex10/links"
)

func crawl(url string, ctx context.Context) []string {
	fmt.Println(url)

	list, err := links.Extract(url, ctx)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	go func() {
		os.Stdin.Read(make([]byte, 1))
		cancel()
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, ctx)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
