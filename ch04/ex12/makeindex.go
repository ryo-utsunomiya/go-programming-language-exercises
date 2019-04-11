package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch04/ex12/xkcd"
)

func main() {
	comics := make([]*xkcd.Comic, 0)
	failCnt := 0

	for i := 1; ; i++ {
		comic, err := fetchComic(i)
		if err != nil {
			fmt.Println(err)
			failCnt++
			if failCnt > 1 {
				break
			}
		}
		if comic != nil {
			comics = append(comics, comic)
		}
		time.Sleep(100)
	}

	if err := makeIndex(comics); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func jsonUrl(id int) string {
	return fmt.Sprintf("https://xkcd.com/%d/info.0.json", id)
}

func fetchComic(id int) (*xkcd.Comic, error) {
	fmt.Printf("Get %s\n", jsonUrl(id))
	resp, err := http.Get(jsonUrl(id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	var result xkcd.Comic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func makeIndex(comics []*xkcd.Comic) error {
	file, err := os.Create("xkcdindex.json")
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(comics); err != nil {
		return err
	}

	return nil
}
