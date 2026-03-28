// Standard Cobra main.go: minimal entrypoint
// Delegates all CLI logic/parsing to cmd/root.go
package main

import (
	"os"

	"GoodFirstGo/cmd"
)

func main() {
	// Execute root command: parses os.Args, binds flags, runs logic
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
