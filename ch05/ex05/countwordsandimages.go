package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countwordsandimages: %v\n", err)
			continue
		}
		fmt.Printf("words: %d, images: %d\n", words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		s := bufio.NewScanner(strings.NewReader(n.Data))
		s.Split(bufio.ScanWords)
		for s.Scan() {
			words++
		}
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}

	return
}
