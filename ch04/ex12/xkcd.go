package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch04/ex12/xkcd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: xkcd {keyword}")
		os.Exit(1)
	}
	keyword := os.Args[1]

	file, err := os.Open("xkcdindex.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	comics := make([]*xkcd.Comic, 0)
	if err := json.NewDecoder(file).Decode(&comics); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := file.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, comic := range comics {
		if comic.Contains(keyword) {
			fmt.Printf("URL: %s\n", comic.Url())
			fmt.Printf("Transcript: %s\n", comic.Transcript)
			fmt.Println()
		}
	}
}
