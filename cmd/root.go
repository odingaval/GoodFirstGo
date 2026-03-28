package cmd

import (
	"fmt"
	"os"

	"GoodFirstGo/internal/github"

	"github.com/spf13/cobra"
)

// Global flags - Cobra binds CLI flags to these vars automatically via Flags().Var()
var (
	limit           int    // --limit: max issues to fetch (default 5)
	language, label string // --language, --label for GitHub search query
)

var rootCmd = &cobra.Command{
	Use:   "goodfirstgo",
	Short: "Find good first issues on GitHub",
	Long:  `CLI tool to search GitHub repositories for good-first-issue labels by language.`,
	// Run: executes when command invoked without subcommands (our root behavior)
	Run: func(cmd *cobra.Command, args []string) {
		// Build GitHub search query using bound flag values
		query := fmt.Sprintf(`label:"%s" language:%s`, label, language)
		fmt.Printf("Fetching %d %s issues with label '%s'...\n", limit, language, label)

		client := github.NewClient()
		issues, err := client.SearchIssues(query, limit)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error fetching issues:", err)
			os.Exit(1)
		}

		fmt.Printf("Found %d matching issues:\n", len(issues))
		for _, issue := range issues {
			fmt.Printf("- %s (Repo: %s) [%s]\n", issue.Title, issue.Repo.FullName, issue.HTMLURL)
		}
	},
}

// Execute is the entrypoint called from main.go - starts Cobra's command parsing/execution
func Execute() error {
	return rootCmd.Execute() // Parses flags/args, validates, runs matching Run func
}

// init() runs automatically: registers flags on rootCmd
func init() {
	// Bind persistent flags to vars (available to root/subcommands)
	rootCmd.Flags().IntVar(&limit, "limit", 5, "Number of issues to fetch")
	rootCmd.Flags().StringVar(&language, "language", "go", "Language to search for")
	rootCmd.Flags().StringVar(&label, "label", "good-first-issue", "Label to search for")
	// Flags have defaults; no required flags needed
}
