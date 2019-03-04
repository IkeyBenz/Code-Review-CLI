package main

import (
	"os"

	"github.com/IkeyBenz/CodeReviewCLI/commands"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "gifm",
		Short:        "Hello Gophers!",
		SilenceUsage: true,
	}

	cmd.AddCommand(commands.PrintTime())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
