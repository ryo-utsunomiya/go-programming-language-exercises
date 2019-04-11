package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: poster {title}")
		os.Exit(1)
	}

	poster, err := fetchPoster(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := downloadPoster(poster); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func buildUrl(title string) string {
	return fmt.Sprintf(
		"https://www.omdbapi.com/?apikey=%s&t=%s",
		os.Getenv("OMDB_APIKEY"),
		url.QueryEscape(title),
	)
}

type Poster struct {
	Title  string
	Poster string
}

func fetchPoster(title string) (*Poster, error) {
	resp, err := http.Get(buildUrl(title))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	poster := Poster{}
	if err := json.NewDecoder(resp.Body).Decode(&poster); err != nil {
		return nil, err
	}

	return &poster, nil
}

func downloadPoster(poster *Poster) error {
	resp, err := http.Get(poster.Poster)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	name := fmt.Sprintf("%s.jpg", poster.Title)
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return err
	}

	return nil
}
