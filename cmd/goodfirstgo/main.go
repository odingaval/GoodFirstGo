package main

import (
	"flag"
	"fmt"

	"GoodFirstGo/internal/github"
)

func main() {
	limit := flag.Int("limit", 5, "Number of issue to fetch") // this will give a refrence to the int which will later hold the parsed value

	flag.Parse() // updated the parsed value at *int
	fmt.Printf("Fetchng %d good first issues... \n", *limit)

	client := github.NewClient()

	issues, err := client.SearchGoodFirstIssues()
	if err != nil {
		fmt.Println("Error fetching issues:", err)
		return
	}

	for i, issue := range issues {
		fmt.Printf("%d. %s\n", i+1, issue.Title)
		fmt.Println(issue.URL)
		fmt.Println()
	}
}
