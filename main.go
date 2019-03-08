package main

import (
	"os"

	"github.com/IkeyBenz/CodeReviewCLI/commands"
)

func main() {
	if err := commands.RootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
