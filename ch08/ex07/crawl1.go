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

var tokens = make(chan struct{}, 10)
var originHostnames = make(map[string]bool)

func crawl(url string) []string {
	fmt.Println(url)

	if err := save(url); err != nil {
		tokens <- struct{}{}
		return nil
	}

	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	for _, arg := range os.Args[1:] {
		u, err := url.Parse(arg)
		if err != nil {
			continue
		}
		originHostnames[u.Hostname()] = true
	}
	fmt.Println(originHostnames)

	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++

				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func save(rawurl string) error {
	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}

	if _, ok := originHostnames[u.Hostname()]; !ok {
		return fmt.Errorf("%s is not one of origin host names", u.Hostname())
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
