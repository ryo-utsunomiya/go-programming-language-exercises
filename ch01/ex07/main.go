package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

func Fetch(client HttpClient, url string, out io.Writer) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(out, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch: copying %s: %v\n", url, err)
		os.Exit(1)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		Fetch(http.DefaultClient, url, os.Stdout)
	}
}
