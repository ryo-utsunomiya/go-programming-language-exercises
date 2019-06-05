package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	mirroredQuery(os.Args[1:])
}

func mirroredQuery(urls []string) string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	responses := make(chan string, 3)
	for _, url := range urls {
		go func() {
			f, _, err := fetch(url, ctx)
			if err != nil {
				log.Print(err)
			}
			cancel()
			responses <- f
		}()
	}
	return <-responses
}

func fetch(url string, ctx context.Context) (filename string, n int64, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", 0, err
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if len(local) == 0 || local == "." || local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)

	if err := f.Close(); err != nil {
		return "", 0, err
	}

	return local, n, err
}
