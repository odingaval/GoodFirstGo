package main

import (
	"flag"
	"fmt"

	"GoodFirstGo/internal/github"
)

func main() {
	limit := flag.Int("limit", 5, "Number of issues to fetch")
	language := flag.String("language", "go", "Language to search for (e.g., 'go', 'javascript')")
	label := flag.String("label", "good-first-issue", "Label to search for")

	flag.Parse()

	query := fmt.Sprintf(`label:"%s" language:%s`, *label, *language)
	fmt.Printf("Fetching %d %s issues with label '%s'...\n", *limit, *language, *label)

	client := github.NewClient()

	issues, err := client.SearchIssues(query, *limit)
	if err != nil {
		fmt.Println("Error fetching issues:", err)
		return
	}

	fmt.Printf("Found %d matching issues:\n", len(issues))
	for _, issue := range issues {
		fmt.Printf("- %s (Repo: %s) [%s]\n", issue.Title, issue.Repo.FullName, issue.HTMLURL)
	}
}
