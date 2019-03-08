package commands

import "github.com/spf13/cobra"

// RootCommand is the main command that the children command use
var RootCommand = &cobra.Command{
	Use:          "code-review",
	Short:        "Code Reviewer:",
	SilenceUsage: true,
}

func init() {
	RootCommand.AddCommand(SignUp)
	RootCommand.AddCommand(SignIn)
	RootCommand.AddCommand(MakeRequest)
	RootCommand.AddCommand(ShowStatus)
}
