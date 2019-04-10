package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("usage: createissue owner repo title content")
		os.Exit(1)
	}

	values := url.Values{}
	values.Set("title", os.Args[3])
	values.Set("body", os.Args[4])

	fmt.Printf("POST %s\n", buildUrl(os.Args[1], os.Args[2]))

	req, err := http.NewRequest(
		"POST",
		buildUrl(os.Args[1], os.Args[2]),
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(
		os.Getenv("GITHUB_USERNAME"),
		os.Getenv("GITHUB_PASSWORD"),
	)

	resp, err := http.Post(
		buildUrl(os.Args[1], os.Args[2]),
		"application/x-www-form-urlencoded",
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("status: %v\n", resp.Status)
	fmt.Printf("header: %v\n", resp.Header)
	fmt.Printf("body: %v\n", string(b))
}

func buildUrl(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
}
