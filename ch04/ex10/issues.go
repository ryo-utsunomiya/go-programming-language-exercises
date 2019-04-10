package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch04/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total %d issues:\n", result.TotalCount)
	fmt.Println()

	fmt.Printf("%d issues created in a month\n", len(result.CreatedInAMonth()))
	for _, item := range result.CreatedInAMonth() {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println()

	fmt.Printf("%d issues created in a year\n", len(result.CreatedInAYear()))
	for _, item := range result.CreatedInAYear() {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println()

	fmt.Printf("%d issues created one or more year ago\n", len(result.CreatedAYearAgo()))
	for _, item := range result.CreatedAYearAgo() {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
