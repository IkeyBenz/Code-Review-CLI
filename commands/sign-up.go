package commands

import (
	"fmt"

	"github.com/IkeyBenz/CodeReviewCLI/database"
	"github.com/IkeyBenz/CodeReviewCLI/resources/users"
	"github.com/spf13/cobra"
)

// SignUp will print the current time in a pretty way
var SignUp = &cobra.Command{
	Use:   "register",
	Short: "Register for code review",
	Long:  "Promps you for an email and a password and creates an account for you on code review.\n\tUsage:code-review register",
	RunE: func(cmd *cobra.Command, args []string) error {
		defer database.MgoSession.Close()
		err := users.NewUser(args[0], args[1])
		if err != nil {
			return err
		}

		fmt.Printf("You are now registered on code-review as %s\n", args[0])

		return nil
	},
}

func hashPassword(p string) (string, error) {
	return "hashedPassword", nil
}
