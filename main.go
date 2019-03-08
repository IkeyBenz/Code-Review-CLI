package main

import (
	"os"

	"github.com/IkeyBenz/CodeReviewCLI/newcommands"
)

func main() {
	if err := newcommands.RootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
