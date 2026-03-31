package output

import (
	"fmt"

	"GoodFirstGo/internal/models"
)

func PrintIssues(issues []models.Issue) {
	fmt.Printf("Found %d good first issues:\n", len(issues))
	for _, issue := range issues {
		fmt.Printf("- %s (Repo: %s) [%s]\n", issue.Title, issue.Repo.FullName, issue.HTMLURL)
		if len(issue.Resources) > 0 {
			fmt.Print("  📚 Learn ")
			for i, res := range issue.Resources {
				if i > 0 {
					fmt.Print(" | ")
				}
				fmt.Print(res)
			}
			fmt.Println()
		}
	}
}
