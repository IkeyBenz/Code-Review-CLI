package commands

import "github.com/spf13/cobra"

// RootCommand is the main command that the children command use
var RootCommand = &cobra.Command{
	Use:          "code-review",
	Short:        "Code Reviewer:",
	Long:         "Code review allows you to send requests for code-review from your peers",
	SilenceUsage: true,
}

func init() {
	RootCommand.AddCommand(SignUp)
	RootCommand.AddCommand(SignIn)
	RootCommand.AddCommand(MakeRequest)
	RootCommand.AddCommand(ShowStatus)
}
