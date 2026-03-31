package cmd

import (
	"fmt"
	"os"

	"github.com/odingaval/GoodFirstGo/internal/github"
	"github.com/odingaval/GoodFirstGo/internal/resources"
	"github.com/odingaval/GoodFirstGo/pkg/output"

	"github.com/spf13/cobra"
)

// Global flags
var (
	limit    int
	language string
	label    string
	learning bool
	stars    int
	age      string
)

var rootCmd = &cobra.Command{
	Use:   "goodfirstgo",
	Short: "Find good first issues on GitHub",
	Long:  `CLI tool to search GitHub repositories for good-first-issue labels by language.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := fmt.Sprintf(`label:"%s" language:%s`, label, language)
		if stars > 0 {
			query += fmt.Sprintf(` stars:>%d`, stars)
		}
		if age != "" {
			// GitHub qualifiers: https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests#search-based-on-when-a-issue-or-pull-request-was-created-or-last-updated
			ageQual := ""
			switch age {
			case "recent":
				ageQual = "created:>=2024-07-01"
			case "week":
				ageQual = "created:>=2024-07-22"
			case "month":
				ageQual = "created:>=2024-06-23"
			}
			if ageQual != "" {
				query += " " + ageQual
			}
		}

		fmt.Printf("Query: %q\n", query)
		fmt.Printf("Fetching %d %s issues...\n", limit, language)

		client := github.NewClient()
		issues, err := client.SearchIssues(query, limit)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		if learning {
			for i := range issues {
				issues[i].Resources = resources.GetLearningResources(language)
			}
			fmt.Println("💡 Learning mode enabled")
		}

		output.PrintIssues(issues)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().IntVar(&limit, "limit", 5, "Max issues")
	rootCmd.Flags().StringVar(&language, "language", "go", "Language")
	rootCmd.Flags().StringVar(&label, "label", "good-first-issue", "Label")
	rootCmd.Flags().BoolVar(&learning, "learning", false, "Show resources")
	rootCmd.Flags().IntVar(&stars, "stars", 0, "Min repo stars")
	rootCmd.Flags().StringVar(&age, "age", "", "Issue age: recent/week/month")
}
