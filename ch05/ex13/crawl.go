package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch05/ex13/links"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

var originHostname string

func crawl(rawurl string) []string {
	fmt.Println(rawurl)

	if err := save(rawurl); err != nil {
		log.Print(err)
	}

	list, err := links.Extract(rawurl)
	if err != nil {
		log.Print(err)
	}
	return list
}

func save(rawurl string) error {
	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}

	if originHostname == "" {
		originHostname = u.Hostname()
	} else if originHostname != u.Hostname() {
		return nil
	}

	if err := os.Mkdir(dirname(u), 0777); err != nil {
		return err
	}

	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename(u))
	if err != nil {
		return err
	}

	if _, err := io.Copy(file, resp.Body); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}

func dirname(u *url.URL) string {
	if filepath.Ext(u.Path) == "" {
		return filepath.Join(u.Hostname(), u.Path)
	} else {
		return filepath.Join(u.Hostname(), filepath.Dir(u.Path))
	}
}

func filename(u *url.URL) string {
	if filepath.Ext(u.Path) == "" {
		return filepath.Join(dirname(u), "index.html")
	} else {
		return u.Path
	}
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
